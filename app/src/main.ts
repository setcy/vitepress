import './assets/fonts.css'
import './assets/base.css'
import './assets/utils.css'
import './assets/vars.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
