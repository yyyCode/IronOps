<template>
  <div class="user-management">
    <div class="page-header">
      <h2>用户管理</h2>
      <div class="header-actions">
        <el-button icon="Refresh" circle @click="fetchUsers" />
        <el-button type="primary" icon="Plus">新建用户</el-button>
      </div>
    </div>

    <el-row :gutter="20" class="stat-cards">
      <el-col :span="8">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ users.length }}</div>
            <div class="stat-label">总用户数</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ activeUsers }}</div>
            <div class="stat-label">活跃用户</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">{{ adminUsers }}</div>
            <div class="stat-label">管理员</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card shadow="never" class="table-card">
      <div class="filter-bar">
        <el-input
          v-model="searchQuery"
          placeholder="搜索用户名"
          prefix-icon="Search"
          style="width: 240px"
        />
        <el-select v-model="roleFilter" placeholder="角色筛选" style="width: 120px">
          <el-option label="所有" value="" />
          <el-option label="管理员" value="admin" />
          <el-option label="运维" value="ops" />
          <el-option label="访客" value="viewer" />
        </el-select>
        <div class="filter-actions">
          <el-button type="primary" @click="fetchUsers">搜索</el-button>
          <el-button @click="resetFilters">重置</el-button>
        </div>
      </div>

      <el-table :data="filteredUsers" style="width: 100%" :header-cell-style="{ background: '#f5f7fa' }" v-loading="loading">
        <el-table-column prop="avatar" width="60">
             <template #default="scope">
                 <el-avatar :size="32" style="background-color: #409EFF">{{ scope.row.username ? scope.row.username.charAt(0).toUpperCase() : 'U' }}</el-avatar>
             </template>
        </el-table-column>
        <el-table-column prop="username" label="用户">
             <template #default="scope">
                 <div class="user-info">
                     <span class="username">{{ scope.row.username }}</span>
                     <span class="user-id">ID: {{ scope.row.ID }}</span>
                 </div>
             </template>
        </el-table-column>
        <el-table-column prop="role" label="角色">
            <template #default="scope">
                <el-tag :type="getRoleType(scope.row.role)" effect="light" size="small">{{ scope.row.role }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column prop="CreatedAt" label="注册时间" />
        <el-table-column label="操作" width="280">
          <template #default>
            <el-button type="primary" link size="small">查看</el-button>
            <el-button type="primary" link size="small">编辑</el-button>
            <el-button type="danger" link size="small">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { getUsers } from '@/api/user';

const searchQuery = ref('');
const roleFilter = ref('');
const users = ref([]);
const loading = ref(false);

const fetchUsers = async () => {
    loading.value = true;
    try {
        const data = await getUsers();
        users.value = data;
    } catch (error) {
        console.error(error);
    } finally {
        loading.value = false;
    }
};

const filteredUsers = computed(() => {
    return users.value.filter(user => {
        const matchQuery = !searchQuery.value || user.username.toLowerCase().includes(searchQuery.value.toLowerCase());
        const matchRole = !roleFilter.value || user.role === roleFilter.value;
        return matchQuery && matchRole;
    });
});

const activeUsers = computed(() => users.value.length); // All active as we don't have status field yet
const adminUsers = computed(() => users.value.filter(u => u.role === 'admin').length);

const resetFilters = () => {
    searchQuery.value = '';
    roleFilter.value = '';
}

const getRoleType = (role) => {
    if (role === 'admin') return 'danger';
    if (role === 'ops') return 'warning';
    return 'info';
}

onMounted(() => {
    fetchUsers();
});
</script>

<style scoped>
.user-management {
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

.user-info {
    display: flex;
    flex-direction: column;
}

.username {
    font-weight: 500;
}

.user-id {
    font-size: 12px;
    color: #909399;
}
</style>
