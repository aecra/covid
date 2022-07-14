// main.ts
import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import './css/normalize.css';
import 'element-plus/dist/index.css';
import 'element-plus/theme-chalk/display.css';
import App from './App.vue';
import conf from './conf';
import OAuthService from './utils/OAuthService';

// import components
import Home from './components/Home.vue';
import User from './components/User.vue';
import Help from './components/Help.vue';
import Clause from './components/Clause.vue';

import { createRouter, createWebHashHistory } from 'vue-router';

console.log("before Init OAuthService");
OAuthService.Init(conf.OAuth);
console.log("after Init OAuthService");

const routes = [
  { path: '/', component: Home },
  { path: '/User', component: User },
  { path: '/Help', component: Help },
  { path: '/Clause', component: Clause },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes: [...routes, ...OAuthService.Routes],
});

router.beforeEach(OAuthService.NavigationGuard);

const app = createApp(App);
app.use(router);
app.use(ElementPlus);
app.mount('#app');
