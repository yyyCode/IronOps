<template>
  <div class="role-management">
    <div class="page-header">
      <h2>角色管理</h2>
      <div class="header-actions">
        <el-button icon="Refresh" circle />
        <el-button type="primary" icon="Plus">新建角色</el-button>
      </div>
    </div>

    <el-row :gutter="20" class="stat-cards">
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">1</div>
            <div class="stat-label">总角色数</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">1</div>
            <div class="stat-label">启用角色</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">0</div>
            <div class="stat-label">系统角色</div>
          </div>
        </el-card>
      </el-col>
       <el-col :span="6">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-value">1</div>
            <div class="stat-label">关联用户</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card shadow="never" class="table-card">
      <div class="filter-bar">
        <el-input
          v-model="searchQuery"
          placeholder="搜索角色名称或编码"
          prefix-icon="Search"
          style="width: 240px"
        />
        <el-select v-model="statusFilter" placeholder="状态筛选" style="width: 120px">
          <el-option label="所有" value="" />
          <el-option label="启用" value="active" />
          <el-option label="禁用" value="disabled" />
        </el-select>
         <el-select v-model="typeFilter" placeholder="类型筛选" style="width: 120px">
          <el-option label="所有" value="" />
          <el-option label="自定义角色" value="custom" />
          <el-option label="系统角色" value="system" />
        </el-select>
        <div class="filter-actions">
          <el-button type="primary">搜索</el-button>
          <el-button>重置</el-button>
        </div>
      </div>

      <el-table :data="roles" style="width: 100%" :header-cell-style="{ background: '#f5f7fa' }">
        <el-table-column prop="roleName" label="角色">
             <template #default="scope">
                 <div class="role-info">
                     <span class="role-name">{{ scope.row.name }}</span>
                     <span class="role-code">{{ scope.row.code }}</span>
                 </div>
             </template>
        </el-table-column>
        <el-table-column prop="type" label="类型">
            <template #default="scope">
                <el-tag :type="scope.row.type === '系统角色' ? 'info' : ''" effect="plain" size="small">{{ scope.row.type }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column prop="status" label="状态">
          <template #default="scope">
            <el-switch v-model="scope.row.active" />
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="permCount" label="权限数" align="center" />
        <el-table-column prop="userCount" label="用户数" align="center" />
        <el-table-column prop="createdAt" label="创建时间" />
        <el-table-column label="操作" width="220">
          <template #default>
            <el-button type="primary" link size="small">查看</el-button>
            <el-button type="primary" link size="small">编辑</el-button>
            <el-button type="primary" link size="small">权限</el-button>
            <el-button type="danger" link size="small">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="pagination-container">
          <el-pagination background layout="prev, pager, next" :total="1" />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getRoles } from '@/api/user';

const searchQuery = ref('');
const statusFilter = ref('');
const typeFilter = ref('');
const roles = ref([]);
const loading = ref(false);

const fetchRoles = async () => {
    loading.value = true;
    try {
        const data = await getRoles();
        // Backend returns ["admin", "ops", "viewer"]
        // We map it to the table structure
        roles.value = data.map(role => ({
            name: role.charAt(0).toUpperCase() + role.slice(1),
            code: role,
            type: 'System Role',
            active: true,
            description: `Default ${role} role`,
            permCount: role === 'admin' ? 'All' : (role === 'ops' ? 'Manage' : 'View'),
            userCount: '-', // We don't have this count yet
            createdAt: new Date().toLocaleDateString()
        }));
    } catch (error) {
        console.error("Failed to fetch roles", error);
    } finally {
        loading.value = false;
    }
};

onMounted(() => {
    fetchRoles();
});
</script>

<style scoped>
.role-management {
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
}

.filter-actions {
  margin-left: auto;
}

.role-info {
    display: flex;
    flex-direction: column;
}

.role-name {
    font-weight: 500;
    color: #303133;
}

.role-code {
    font-size: 12px;
    color: #909399;
    background-color: #f4f4f5;
    padding: 2px 4px;
    border-radius: 4px;
    width: fit-content;
    margin-top: 4px;
}

.pagination-container {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
}
</style>
