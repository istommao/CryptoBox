import {createApp} from 'vue'
import App from './App.vue'
import * as VueRouter from 'vue-router';
// import HelloWorldVue from './components/HelloWorld.vue';

import HomePage from './components/HomePage.vue';

import Ed25519Vue from './components/Ed25519.vue';
import X25519Page from './components/X25519.vue';

import 'element-plus/theme-chalk/dark/css-vars.css'


const routes = [
    { path: '/', component: HomePage, name: "Home" },
    { path: '/ed25519', component: Ed25519Vue },
    { path: '/x25519', component: X25519Page },
  ]

  
const router = VueRouter.createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: VueRouter.createWebHashHistory(),
    routes, // short for `routes: routes`
  })


const app = createApp(App)


app.use(router)

app.mount('#app')
