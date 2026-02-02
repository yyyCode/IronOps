<template>
  <div class="alert-channels-container">
    <div class="page-header">
      <h2>通知渠道配置</h2>
      <el-button type="primary" @click="showCreateDialog = true">新建渠道</el-button>
    </div>

    <el-card shadow="hover" class="table-card">
      <el-table :data="channels" style="width: 100%" v-loading="loading">
        <el-table-column prop="ID" label="ID" width="80" />
        <el-table-column prop="name" label="渠道名称" width="180" />
        <el-table-column prop="type" label="类型" width="120">
          <template #default="scope">
            <el-tag>{{ scope.row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="config" label="配置信息">
          <template #default="scope">
            <span class="code-text">{{ truncate(scope.row.config, 50) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="enabled" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.enabled ? 'success' : 'info'" effect="plain">
              {{ scope.row.enabled ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button link type="danger" size="small" @click="handleDelete(scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Create Dialog -->
    <el-dialog v-model="showCreateDialog" title="新建通知渠道" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="渠道名称">
          <el-input v-model="form.name" placeholder="例如：运维组飞书群" />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="form.type" placeholder="选择类型">
            <el-option label="飞书 (Feishu)" value="feishu" />
            <el-option label="钉钉 (DingTalk)" value="dingtalk" />
            <el-option label="通用 Webhook" value="webhook" />
          </el-select>
        </el-form-item>
        <el-form-item label="Webhook URL">
          <el-input v-model="webhookUrl" placeholder="https://..." />
          <div class="form-tip">请输入完整的 Webhook 地址</div>
        </el-form-item>
        <el-form-item label="启用">
          <el-switch v-model="form.enabled" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showCreateDialog = false">取消</el-button>
          <el-button type="primary" @click="handleCreate">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { getAlertChannels, createAlertChannel, deleteAlertChannel } from '@/api/alert';
import { ElMessage, ElMessageBox } from 'element-plus';

const channels = ref([]);
const loading = ref(false);
const showCreateDialog = ref(false);
const webhookUrl = ref('');

const form = reactive({
  name: '',
  type: 'feishu',
  config: '',
  enabled: true
});

const fetchChannels = async () => {
  loading.value = true;
  try {
    const data = await getAlertChannels();
    channels.value = data || [];
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const handleCreate = async () => {
  if (!webhookUrl.value) {
    ElMessage.warning('请输入 Webhook URL');
    return;
  }

  // Pack config into JSON string
  const configObj = { url: webhookUrl.value };
  form.config = JSON.stringify(configObj);

  try {
    await createAlertChannel(form);
    ElMessage.success('创建成功');
    showCreateDialog.value = false;
    fetchChannels();
    // Reset form
    form.name = '';
    webhookUrl.value = '';
  } catch (error) {
    ElMessage.error('创建失败');
  }
};

const handleDelete = (id) => {
  ElMessageBox.confirm('确定要删除该渠道吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    try {
      await deleteAlertChannel(id);
      ElMessage.success('删除成功');
      fetchChannels();
    } catch (error) {
      ElMessage.error('删除失败');
    }
  });
};

const truncate = (str, n) => {
  return (str.length > n) ? str.substr(0, n - 1) + '...' : str;
};

onMounted(() => {
  fetchChannels();
});
</script>

<style scoped>
.alert-channels-container {
  padding: 20px;
}
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
.page-header h2 {
  margin: 0;
  color: #303133;
}
.code-text {
  font-family: monospace;
  color: #606266;
}
.form-tip {
  font-size: 12px;
  color: #909399;
  line-height: 1.5;
  margin-top: 5px;
}
</style>