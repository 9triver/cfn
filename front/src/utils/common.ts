import {nextTick, Ref} from "vue";
import router from '/@/router/index';
import {useTitle} from "@vueuse/core";

/**
 * 根据路由 meta.title 设置浏览器标题
 */
export function setTitleFromRoute() {
    // noinspection JSIgnoredPromiseFromCall
    nextTick(() => {
        console.log(router.currentRoute.value.meta)
        if (typeof router.currentRoute.value.meta.title != 'string') {
            return;
        }
        const webTitle =
            // i18n.global.te(router.currentRoute.value.meta.title)
            // ? i18n.global.t(router.currentRoute.value.meta.title) :
            router.currentRoute.value.meta.title;
        const title = useTitle() as Ref<string>;
        const siteConfig = {
            siteName: "example",
        };
        title.value = `${webTitle}${siteConfig.siteName ? ' - ' + siteConfig.siteName : ''}`;
    });
}
