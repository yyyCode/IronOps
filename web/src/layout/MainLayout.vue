<template>
  <div class="common-layout">
    <el-container class="layout-container">
      <el-aside width="240px" class="aside">
        <div class="logo">
          <el-icon><Monitor /></el-icon>
          <span>IronOps</span>
        </div>
        <el-menu
          active-text-color="#409EFF"
          background-color="#001529"
          class="el-menu-vertical-demo"
          default-active="/dashboard"
          text-color="#fff"
          router
        >
          <el-sub-menu index="1">
            <template #title>
              <el-icon><Setting /></el-icon>
              <span>系统管理</span>
            </template>
            <el-menu-item index="/dashboard">
              <el-icon><Odometer /></el-icon>欢迎页
            </el-menu-item>
            <el-menu-item index="/users">
              <el-icon><User /></el-icon>用户管理
            </el-menu-item>
             <el-menu-item index="/roles">
              <el-icon><Lock /></el-icon>角色权限
            </el-menu-item>
            <el-menu-item index="/audit">
              <el-icon><Document /></el-icon>审计日志
            </el-menu-item>
          </el-sub-menu>

          <el-sub-menu index="2">
            <template #title>
              <el-icon><Connection /></el-icon>
              <span>服务治理</span>
            </template>
            <el-menu-item index="/services">
              <el-icon><Files /></el-icon>服务树
            </el-menu-item>
          </el-sub-menu>
          
           <el-sub-menu index="3">
            <template #title>
              <el-icon><Bell /></el-icon>
              <span>监控告警</span>
            </template>
            <el-menu-item index="/alerts">
              <el-icon><BellFilled /></el-icon>告警列表
            </el-menu-item>
            <el-menu-item index="/alert-rules">
              <el-icon><Operation /></el-icon>告警规则
            </el-menu-item>
            <el-menu-item index="/alert-channels">
              <el-icon><Message /></el-icon>通知渠道
            </el-menu-item>
          </el-sub-menu>

        </el-menu>
      </el-aside>
      <el-container>
        <el-header class="header">
          <div class="header-left">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
              <el-breadcrumb-item>当前页面</el-breadcrumb-item>
            </el-breadcrumb>
          </div>
          <div class="header-right">
             <el-button circle icon="Search" />
             <el-button circle icon="Setting" />
             <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
          </div>
        </el-header>
        <el-main class="main-content">
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.layout-container {
  height: 100vh;
}

.aside {
  background-color: #001529;
  color: #fff;
  transition: width 0.3s;
  display: flex;
  flex-direction: column;
}

.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: bold;
  background-color: #002140;
  gap: 10px;
}

.el-menu {
  border-right: none;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #dcdfe6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 15px;
}

.main-content {
  padding: 20px;
  background-color: #f0f2f5;
}
</style>
