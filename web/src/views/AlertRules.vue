<template>
  <div class="alert-rules-container">
    <div class="page-header">
      <h2>告警规则管理</h2>
      <el-button type="primary" @click="showCreateDialog = true">新建规则</el-button>
    </div>

    <el-card shadow="hover" class="table-card">
      <el-table :data="rules" style="width: 100%" v-loading="loading">
        <el-table-column prop="ID" label="ID" width="80" />
        <el-table-column prop="name" label="规则名称" width="180" />
        <el-table-column prop="metric_type" label="监控指标" width="120">
          <template #default="scope">
            <el-tag>{{ scope.row.metric_type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="触发条件" width="180">
          <template #default="scope">
            {{ scope.row.condition }} {{ scope.row.threshold }}
          </template>
        </el-table-column>
        <el-table-column prop="severity" label="严重程度" width="120">
          <template #default="scope">
            <el-tag :type="getSeverityType(scope.row.severity)">{{ scope.row.severity }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="enabled" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.enabled ? 'success' : 'info'" effect="plain">
              {{ scope.row.enabled ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" />
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button link type="danger" size="small" @click="handleDelete(scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Create Dialog -->
    <el-dialog v-model="showCreateDialog" title="新建告警规则" width="500px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="规则名称">
          <el-input v-model="form.name" placeholder="例如：CPU过高警告" />
        </el-form-item>
        <el-form-item label="监控指标">
          <el-select v-model="form.metric_type" placeholder="选择指标">
            <el-option label="CPU" value="cpu" />
            <el-option label="Memory" value="memory" />
          </el-select>
        </el-form-item>
        <el-form-item label="条件">
          <el-row :gutter="10">
            <el-col :span="8">
              <el-select v-model="form.condition">
                <el-option label="大于" value=">" />
                <el-option label="小于" value="<" />
                <el-option label="等于" value="=" />
              </el-select>
            </el-col>
            <el-col :span="16">
              <el-input-number v-model="form.threshold" :precision="1" :step="1" style="width: 100%" />
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="严重程度">
          <el-select v-model="form.severity">
            <el-option label="Critical" value="critical" />
            <el-option label="Warning" value="warning" />
            <el-option label="Info" value="info" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="form.description" type="textarea" />
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
import { getAlertRules, createAlertRule, deleteAlertRule } from '@/api/alert';
import { ElMessage, ElMessageBox } from 'element-plus';

const rules = ref([]);
const loading = ref(false);
const showCreateDialog = ref(false);

const form = reactive({
  name: '',
  metric_type: 'cpu',
  condition: '>',
  threshold: 80,
  severity: 'warning',
  description: '',
  enabled: true
});

const fetchRules = async () => {
  loading.value = true;
  try {
    const data = await getAlertRules();
    rules.value = data || [];
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const handleCreate = async () => {
  try {
    await createAlertRule(form);
    ElMessage.success('创建成功');
    showCreateDialog.value = false;
    fetchRules();
    // Reset form
    form.name = '';
    form.description = '';
  } catch (error) {
    ElMessage.error('创建失败');
  }
};

const handleDelete = (id) => {
  ElMessageBox.confirm('确定要删除该规则吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    try {
      await deleteAlertRule(id);
      ElMessage.success('删除成功');
      fetchRules();
    } catch (error) {
      ElMessage.error('删除失败');
    }
  });
};

const getSeverityType = (severity) => {
  switch (severity) {
    case 'critical': return 'danger';
    case 'warning': return 'warning';
    default: return 'info';
  }
};

onMounted(() => {
  fetchRules();
});
</script>

<style scoped>
.alert-rules-container {
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
</style>