import './assets/fonts.css'
import './assets/base.css'
import './assets/utils.css'
import './assets/vars.css'
import './assets/components/custom-block.css'
import './assets/components/vp-code.css'
import './assets/components/vp-code-group.css'
import './assets/components/vp-doc.css'
import './assets/components/vp-sponsor.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
