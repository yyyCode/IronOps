<template>
  <div class="alert-list">
    <div class="page-header">
      <h2>告警中心</h2>
      <div class="header-actions">
        <el-button icon="Refresh" circle @click="fetchAlerts" />
        <el-button type="danger" icon="Bell">全部静音</el-button>
      </div>
    </div>

    <el-row :gutter="20" class="stat-cards">
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value error-text">{{ activeAlertsCount }}</div>
            <div class="stat-label">当前告警</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value warning-text">{{ todayAlertsCount }}</div>
            <div class="stat-label">今日新增</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value success-text">{{ resolvedAlertsCount }}</div>
            <div class="stat-label">已恢复</div>
          </div>
        </el-card>
      </el-col>
       <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">99.9%</div>
            <div class="stat-label">健康度</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card shadow="never" class="table-card">
      <div class="filter-bar">
        <el-input
          v-model="searchQuery"
          placeholder="搜索告警名称或对象"
          prefix-icon="Search"
          style="width: 240px"
        />
        <el-select v-model="statusFilter" placeholder="状态" style="width: 120px">
          <el-option label="全部" value="" />
          <el-option label="活动中" value="firing" />
          <el-option label="已恢复" value="resolved" />
        </el-select>
        <div class="filter-actions">
          <el-button type="primary" @click="fetchAlerts">搜索</el-button>
          <el-button @click="resetFilters">重置</el-button>
        </div>
      </div>

      <el-table :data="filteredAlerts" style="width: 100%" :header-cell-style="{ background: '#f5f7fa' }" v-loading="loading">
        <el-table-column type="selection" width="55" />
        <el-table-column prop="type" label="告警类型" min-width="150">
            <template #default="scope">
                <span class="alert-name">{{ scope.row.type }}</span>
            </template>
        </el-table-column>
        <el-table-column prop="instance_id" label="实例ID" width="100" />
        <el-table-column prop="message" label="告警内容" min-width="200" show-overflow-tooltip />
        <el-table-column prop="CreatedAt" label="开始时间" width="180" />
        <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
                <el-tag :type="getStatusType(scope.row.status)" effect="plain" size="small">
                    <span class="status-dot" :class="scope.row.status"></span>
                    {{ getStatusLabel(scope.row.status) }}
                </el-tag>
            </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default>
            <el-button type="primary" link size="small">详情</el-button>
            <el-button type="primary" link size="small">认领</el-button>
            <el-button type="danger" link size="small">关闭</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { getAlerts } from '@/api/alert';

const searchQuery = ref('');
const statusFilter = ref('');
const alerts = ref([]);
const loading = ref(false);

const fetchAlerts = async () => {
    loading.value = true;
    try {
        const data = await getAlerts();
        alerts.value = data;
    } catch (error) {
        console.error(error);
    } finally {
        loading.value = false;
    }
}

const filteredAlerts = computed(() => {
  return alerts.value.filter(alert => {
    const matchesSearch = !searchQuery.value || 
      alert.type.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
      alert.message.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesStatus = !statusFilter.value || alert.status === statusFilter.value;
    
    return matchesSearch && matchesStatus;
  });
});

const activeAlertsCount = computed(() => alerts.value.filter(a => a.status === 'firing').length);
const resolvedAlertsCount = computed(() => alerts.value.filter(a => a.status === 'resolved').length);
const todayAlertsCount = computed(() => {
    const today = new Date().toISOString().split('T')[0];
    return alerts.value.filter(a => a.CreatedAt.startsWith(today)).length;
});

const resetFilters = () => {
  searchQuery.value = '';
  statusFilter.value = '';
};

const getStatusType = (status) => {
  const map = {
    firing: 'danger',
    resolved: 'success',
    silenced: 'info'
  };
  return map[status] || 'info';
};

const getStatusLabel = (status) => {
  const map = {
    firing: '活动中',
    resolved: '已恢复',
    silenced: '已静音'
  };
  return map[status] || status;
};

onMounted(() => {
    fetchAlerts();
});
</script>

<style scoped>
.alert-list {
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-actions {
  display: flex;
  gap: 10px;
}

.stat-cards {
  margin-bottom: 20px;
}

.stat-card {
  height: 100px;
  display: flex;
  align-items: center;
}

.stat-content {
  text-align: center;
  width: 100%;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-top: 5px;
}

.error-text { color: #F56C6C; }
.warning-text { color: #E6A23C; }
.success-text { color: #67C23A; }

.table-card {
  margin-top: 20px;
}

.filter-bar {
  display: flex;
  gap: 15px;
  margin-bottom: 20px;
  align-items: center;
}

.filter-actions {
  margin-left: auto;
}

.alert-name {
    font-weight: 500;
}

.status-dot {
    display: inline-block;
    width: 6px;
    height: 6px;
    border-radius: 50%;
    margin-right: 5px;
    background-color: #909399;
}

.status-dot.firing { background-color: #F56C6C; }
.status-dot.resolved { background-color: #67C23A; }
.status-dot.silenced { background-color: #909399; }
</style>
