<template>
  <div class="audit-log">
    <div class="page-header">
      <h2>审计日志</h2>
      <div class="header-actions">
        <el-button icon="Refresh">刷新</el-button>
        <el-button icon="Download">导出</el-button>
      </div>
    </div>

    <el-row :gutter="20" class="stat-cards">
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ totalLogs }}</div>
            <div class="stat-label">总日志数</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ todayLogs }}</div>
            <div class="stat-label">今日新增</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value error-text">{{ errorLogs }}</div>
            <div class="stat-label">错误日志</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ uniqueUsers }}</div>
            <div class="stat-label">独立用户数</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card shadow="never" class="table-card">
      <div class="filter-bar">
        <el-input
          v-model="searchQuery"
          placeholder="搜索..."
          prefix-icon="Search"
          style="width: 180px"
        />
        <el-select v-model="typeFilter" placeholder="操作类型" style="width: 120px">
          <el-option label="CREATE" value="CREATE" />
          <el-option label="UPDATE" value="UPDATE" />
          <el-option label="DELETE" value="DELETE" />
        </el-select>
        <el-select v-model="targetFilter" placeholder="目标类型" style="width: 120px">
           <el-option label="User" value="User" />
           <el-option label="Role" value="Role" />
        </el-select>
         <el-select v-model="statusFilter" placeholder="状态码" style="width: 120px">
           <el-option label="200" value="200" />
           <el-option label="500" value="500" />
        </el-select>
        <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="-"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            style="width: 240px"
        />
        <div class="filter-actions">
          <el-button type="primary">搜索</el-button>
          <el-button>重置</el-button>
          <el-button link>高级搜索</el-button>
        </div>
      </div>

      <el-table :data="paginatedLogs" v-loading="loading" style="width: 100%" :header-cell-style="{ background: '#f5f7fa' }">
        <el-table-column type="selection" width="55" />
        <el-table-column prop="created_at" label="时间" width="180">
             <template #default="scope">
                 {{ new Date(scope.row.created_at).toLocaleString() }}
             </template>
        </el-table-column>
        <el-table-column prop="user" label="用户" width="180">
             <template #default="scope">
                 <div class="user-info">
                     <span class="username">{{ scope.row.user }}</span>
                 </div>
             </template>
        </el-table-column>
        <el-table-column prop="action" label="操作类型" width="100">
          <template #default="scope">
            <el-tag :type="getActionType(scope.row.action)" effect="plain" size="small">{{ scope.row.action }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="detail" label="详情">
             <template #default="scope">
                 <div class="request-info">
                     <span class="path">{{ scope.row.detail }}</span>
                 </div>
             </template>
        </el-table-column>
        <el-table-column prop="result" label="结果" width="80">
            <template #default="scope">
                <el-tag :type="scope.row.result === 'success' ? 'success' : 'danger'" effect="light" size="small">{{ scope.row.result }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column prop="target" label="目标" width="140" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default>
            <el-button type="primary" link size="small">查看</el-button>
            <el-button type="danger" link size="small">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
       <div class="pagination-container">
          <el-pagination 
            background 
            layout="total, prev, pager, next" 
            :total="filteredAuditLogs.length" 
            :page-size="pageSize"
            v-model:current-page="currentPage"
            @current-change="handlePageChange"
          />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { getAuditLogs } from '@/api/audit';

const searchQuery = ref('');
const typeFilter = ref('');
const targetFilter = ref('');
const statusFilter = ref('');
const dateRange = ref('');
const loading = ref(false);

const currentPage = ref(1);
const pageSize = ref(10);

const auditLogs = ref([]);

const totalLogs = computed(() => auditLogs.value.length);

const todayLogs = computed(() => {
  const today = new Date().toISOString().split('T')[0];
  return auditLogs.value.filter(log => log.created_at && log.created_at.startsWith(today)).length;
});

const errorLogs = computed(() => auditLogs.value.filter(log => log.result === 'fail').length);

const uniqueUsers = computed(() => {
    const users = auditLogs.value.map(log => log.user).filter(u => u);
    return new Set(users).size;
});

const fetchAuditLogs = async () => {
  loading.value = true;
  try {
    const data = await getAuditLogs();
    auditLogs.value = data || [];
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const filteredAuditLogs = computed(() => {
  return auditLogs.value.filter(log => {
    // Basic filtering logic
    const matchSearch = !searchQuery.value || 
      (log.user && log.user.toLowerCase().includes(searchQuery.value.toLowerCase())) ||
      (log.target && log.target.toLowerCase().includes(searchQuery.value.toLowerCase()));
    
    const matchType = !typeFilter.value || log.action === typeFilter.value;
    const matchTarget = !targetFilter.value || (log.target && log.target.includes(targetFilter.value)); // Simple partial match for target
    
    // Check status code (result field in DB is 'success' or 'fail', or status code if we stored it. 
    // The current model has 'result' string. Let's map it roughly or assume statusFilter is for result)
    // If statusFilter is '200', we look for 'success'. If '500', 'fail'.
    let matchStatus = true;
    if (statusFilter.value) {
        if (statusFilter.value === '200') matchStatus = log.result === 'success';
        if (statusFilter.value === '500') matchStatus = log.result === 'fail';
    }

    return matchSearch && matchType && matchTarget && matchStatus;
  });
});

const paginatedLogs = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value;
  const end = start + pageSize.value;
  return filteredAuditLogs.value.slice(start, end);
});

const handlePageChange = (val) => {
  currentPage.value = val;
};

const getActionType = (action) => {
  switch (action) {
    case 'POST': return 'success'; // Create
    case 'PUT': return 'primary'; // Update
    case 'DELETE': return 'danger'; // Delete
    case 'GET': return 'info';
    default: return 'info';
  }
};

onMounted(() => {
  fetchAuditLogs();
});
</script>

<style scoped>
.audit-log {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  padding: 15px 20px;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
}

.page-header h2 {
  margin: 0;
  font-size: 18px;
  color: #303133;
}

.stat-cards {
  margin-bottom: 0;
}

.stat-card {
  border: none;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
}

.stat-content {
  text-align: center;
  padding: 10px 0;
}

.stat-value {
  font-size: 32px;
  font-weight: 600;
  color: #409EFF;
  margin-bottom: 5px;
}

.error-text {
    color: #F56C6C;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.table-card {
  border: none;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
}

.filter-bar {
  display: flex;
  gap: 15px;
  margin-bottom: 20px;
  align-items: center;
  flex-wrap: wrap;
}

.filter-actions {
  margin-left: auto;
}

.user-info {
    font-weight: 500;
    color: #303133;
}

.request-info {
    display: flex;
    flex-direction: column;
    gap: 4px;
}

.method-tag {
    width: fit-content;
}

.path {
    font-family: monospace;
    font-size: 12px;
    color: #606266;
}

.duration-text {
    font-weight: 600;
}

.pagination-container {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
}
</style>
