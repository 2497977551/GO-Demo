import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

import Index from '../components/admin/Index.vue'
import UserList from '../components/user/UserList.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin,
    children: [
      { path: 'index', component: Index },
      { path: 'userlist', component: UserList },
      { path: 'addart', component: AddArt },
      { path: 'artlist', component: ArtList },
      { path: 'catelist', component: CateList }
    ]
  }
]

const router = new VueRouter({
  routes
})
router.beforeEach((to, from, next) => {
  const token = window.localStorage.getItem('Token')
  if (to.path === '/login') {
    return next()
  }
  if (!token && to.path === '/admin') {
    next('/login')
  } else {
    next()
  }
})
export default router
