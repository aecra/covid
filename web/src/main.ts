// main.ts
import {createApp} from 'vue';
import ElementPlus from 'element-plus';
import './css/normalize.css';
import 'element-plus/dist/index.css';
import 'element-plus/theme-chalk/display.css';
import App from './App.vue';

// import components
import LoginPage from './components/Login.vue';
import RegisterPage from './components/Register.vue';
import Home from './components/Home.vue';
import User from './components/User.vue';
import Records from './components/Records.vue'
import Help from './components/Help.vue';
import Clause from './components/Clause.vue';

import {createRouter, createWebHashHistory} from 'vue-router';
import auth from "./utils/auth";

const routes = [
  {
    path: '/login',
    name: 'loginPage',
    component: LoginPage,
    meta: {
      keepalive: false
    }
  },
  {
    path: '/register',
    name: 'registerPage',
    component: RegisterPage,
    meta: {
      keepalive: false
    }
  },
  {path: '/', component: Home, meta: {keepalive: true}},
  {path: '/User', component: User, meta: {keepalive: true}},
  {path: '/Records', component: Records, meta: {keepalive: true}},
  {path: '/Help', component: Help, meta: {keepalive: true}},
  {path: '/Clause', component: Clause, meta: {keepalive: true}},
];

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

router.beforeEach(auth.guard);

const app = createApp(App);
app.use(router);
app.use(ElementPlus);
app.mount('#app');
