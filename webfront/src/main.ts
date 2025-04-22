import { createApp } from 'vue';
import App from './App.vue';
import router from '@/router';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import mitt from 'mitt';

// modules import mark, Please do not remove.

async function start() {
    const app = createApp(App);
    // app.use(pinia);

    app.use(router);
    app.use(Antd);

    app.mount('#app');

    // modules start mark, Please do not remove.

    app.config.globalProperties.eventBus = mitt();
}

// noinspection JSIgnoredPromiseFromCall
start();
