import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import IndexView from '../views/IndexView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import UploadView from '../views/UploadView.vue'
import AllBooksView from '../views/AllBooksView.vue'
import MyBooksView from '../views/MyBooksView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'index',    //name是path的简写，或者说另一个名字，有通过name的跳转
    component: IndexView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: '/upload',
    name: 'upload',
    component: UploadView
  },
  {
    path: '/all-books',
    name: 'all-books',
    component: AllBooksView
  },
  {
    path: '/my-books',
    name: 'my-books',
    component: MyBooksView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach(async (to, from) => {
  let isAuthenticated = true
  let needLogin = false
  if (to.name !== 'register' && to.name !== 'all-books') {
    needLogin =true
  }
  if (localStorage.getItem("userInfo") === null) {
    isAuthenticated = false
  } 
  if (
    // 检查用户是否已登录
    !isAuthenticated &&
    // ❗️ 避免无限重定向
    to.name !== 'login' &&
    // 注册页面不需要登录
    needLogin
  ) {
    // 将用户重定向到登录页面
    return { name: 'login' }
  }
})

export default router
