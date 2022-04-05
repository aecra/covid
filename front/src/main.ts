// main.ts
import { createApp, h } from 'vue';
import ElementPlus from 'element-plus';
import './css/normalize.css';
import 'element-plus/dist/index.css';
import 'element-plus/theme-chalk/display.css'
import App from './App.vue';
import Home from './components/Home.vue';
import User from './components/User.vue';
import Help from './components/Help.vue';
import Clause from './components/Clause.vue';
import Oauth from './components/Oauth.vue';

import {createRouter,createWebHashHistory } from 'vue-router';

const routes = [
  { path: '/', component: Home },
  { path: '/User', component: User },
  { path: '/Help', component: Help },
  { path: '/Oauth', component: Oauth },
  {path: '/Clause', component: Clause}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

const app = createApp(App);
app.use(router);
app.use(ElementPlus);
app.mount('#app');
