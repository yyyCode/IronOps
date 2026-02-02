<template>
  <div class="service-list">
    <div class="toolbar">
      <el-button type="primary" icon="Plus" @click="dialogVisible = true">新增服务</el-button>
      <el-input
        v-model="searchQuery"
        placeholder="搜索服务名称"
        prefix-icon="Search"
        style="width: 200px; margin-left: 10px;"
      />
      <el-button icon="Refresh" circle style="margin-left: 10px" @click="fetchServices" />
    </div>

    <el-table :data="services" style="width: 100%; margin-top: 20px;" border v-loading="loading">
      <el-table-column type="expand">
        <template #default="props">
          <div style="padding: 10px 20px;">
            <div style="margin-bottom: 10px; display: flex; justify-content: space-between; align-items: center;">
               <h4>实例列表</h4>
               <el-button size="small" type="primary" plain icon="Plus" @click="openInstanceDialog(props.row.ID)">添加实例</el-button>
            </div>
           
            <el-table :data="props.row.instances" border size="small">
              <el-table-column label="实例ID" prop="ID" width="80" />
              <el-table-column label="IP地址" prop="ip" />
              <el-table-column label="端口" prop="port" width="100" />
              <el-table-column label="状态" width="100">
                <template #default="scope">
                  <el-tag :type="scope.row.status === 'running' ? 'success' : 'danger'">
                    {{ scope.row.status || 'unknown' }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column label="最后心跳" prop="last_heartbeat" />
              <el-table-column label="操作">
                 <template #default="scope">
                    <el-button size="small" type="danger" link @click="handleAction(scope.row, 'stop')">停止</el-button>
                    <el-button size="small" type="primary" link @click="handleAction(scope.row, 'restart')">重启</el-button>
                 </template>
              </el-table-column>
            </el-table>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="服务名称" prop="name" />
      <el-table-column label="负责人" prop="owner" />
      <el-table-column label="环境" prop="env">
        <template #default="scope">
          <el-tag>{{ scope.row.env }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="实例数">
        <template #default="scope">
          {{ scope.row.instances ? scope.row.instances.length : 0 }}
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="CreatedAt" />
    </el-table>

    <!-- Create Service Dialog -->
    <el-dialog v-model="dialogVisible" title="新增服务" width="30%">
      <el-form :model="form" label-width="80px">
        <el-form-item label="服务名称">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="负责人">
          <el-input v-model="form.owner" />
        </el-form-item>
        <el-form-item label="环境">
          <el-select v-model="form.env" placeholder="Select">
            <el-option label="Dev" value="dev" />
            <el-option label="Test" value="test" />
            <el-option label="Prod" value="prod" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleAddService">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- Add Instance Dialog -->
    <el-dialog v-model="instanceDialogVisible" title="添加实例" width="30%">
      <el-form :model="instanceForm" label-width="80px">
        <el-form-item label="IP地址">
          <el-input v-model="instanceForm.ip" />
        </el-form-item>
        <el-form-item label="端口">
          <el-input v-model.number="instanceForm.port" type="number" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="instanceDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleAddInstance">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { getServices, createService, addInstance, controlInstance } from '@/api/service';

const searchQuery = ref('');
const dialogVisible = ref(false);
const instanceDialogVisible = ref(false);
const loading = ref(false);
const services = ref([]);

const form = ref({
  name: '',
  owner: '',
  env: 'dev'
});

const instanceForm = ref({
    service_id: 0,
    ip: '',
    port: 8080,
    status: 'stopped'
});

const fetchServices = async () => {
    loading.value = true;
    try {
        const data = await getServices();
        services.value = data;
    } catch (error) {
        console.error(error);
    } finally {
        loading.value = false;
    }
}

onMounted(() => {
    fetchServices();
});

const handleAddService = async () => {
  try {
      await createService(form.value);
      ElMessage.success('服务创建成功');
      dialogVisible.value = false;
      fetchServices();
      // Reset form
      form.value = { name: '', owner: '', env: 'dev' };
  } catch (error) {
      // Error handled in request.js
  }
};

const openInstanceDialog = (serviceId) => {
    instanceForm.value.service_id = serviceId;
    instanceForm.value.ip = '';
    instanceForm.value.port = 8080;
    instanceDialogVisible.value = true;
}

const handleAddInstance = async () => {
    try {
        await addInstance(instanceForm.value);
        ElMessage.success('实例添加成功');
        instanceDialogVisible.value = false;
        fetchServices();
    } catch (error) {
        console.error(error);
    }
}

const handleAction = async (instance, action) => {
    try {
        await controlInstance(instance.ID, action);
        ElMessage.success(`操作成功: ${action}`);
        fetchServices(); // Refresh status
    } catch (error) {
        console.error(error);
    }
}
</script>

<style scoped>
.toolbar {
  display: flex;
  align-items: center;
}
</style>
