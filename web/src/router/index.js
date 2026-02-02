import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from '../layout/MainLayout.vue'
import Dashboard from '../views/Dashboard.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: MainLayout,
      redirect: '/dashboard',
      children: [
        {
          path: 'dashboard',
          name: 'dashboard',
          component: Dashboard
        },
        {
          path: 'users',
          name: 'users',
          component: () => import('../views/UserManagement.vue')
        },
        {
          path: 'roles',
          name: 'roles',
          component: () => import('../views/RoleManagement.vue')
        },
        {
          path: 'services',
          name: 'services',
          component: () => import('../views/ServiceList.vue')
        },
        {
          path: 'audit',
          name: 'audit',
          component: () => import('../views/AuditLog.vue')
        },
        {
          path: 'alerts',
          name: 'alerts',
          component: () => import('../views/AlertList.vue')
        },
        {
          path: 'alert-rules',
          name: 'alert-rules',
          component: () => import('../views/AlertRules.vue')
        },
        {
          path: 'alert-channels',
          name: 'alert-channels',
          component: () => import('../views/AlertChannels.vue')
        }
      ]
    }
  ]
})

export default router
