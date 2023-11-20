import {computed, type ComputedRef, ref, type Ref} from 'vue'
import axios from "axios";
import {apiUrl} from "@/support/shared";
import {useMediaQuery} from "@vueuse/core";
import {useRoute} from "vue-router";

export interface Sidebar {
    text?: string
    link?: string
    items?: Sidebar[]
}

export interface SidebarControl {
    isOpen: Ref<boolean>
    isSidebarEnabled: ComputedRef<boolean>
    open: () => void
    close: () => void
    toggle: () => void
}

export interface SidebarItemControl {
    collapsed: Ref<boolean>
    isLink: ComputedRef<boolean>
    isActiveLink: Ref<boolean>
    hasActiveLink: ComputedRef<boolean>
    hasChildren: ComputedRef<boolean>
    toggle: () => void
}

export function useSidebar(Loading: Ref<boolean>, sidebarGroups: Ref<Sidebar[]>) {
    axios.get(apiUrl + "/_meta").then((res) => {
        sidebarGroups.value = res.data.data;
        Loading.value = false;
    });
}

export function useSidebarControl(): SidebarControl {
    const is960 = useMediaQuery('(min-width: 960px)')

    const isOpen = ref(false)

    const isSidebarEnabled = computed(() => is960.value)

    function open() {
        isOpen.value = true
    }

    function close() {
        isOpen.value = false
    }

    function toggle() {
        isOpen.value ? close() : open()
    }

    return {
        isOpen,
        isSidebarEnabled,
        open,
        close,
        toggle
    }
}

export function useSidebarItemControl(
    item: ComputedRef<Sidebar>
): {
    isLink: ComputedRef<boolean>;
    isActiveLink: Ref<boolean>;
    hasActiveLink: ComputedRef<boolean>;
    collapsed: Ref<boolean>;
    hasChildren: ComputedRef<boolean>;
    toggle: () => void
} {

    const collapsed = ref(false)

    const isLink = computed(() => {
        return !!item.value.link
    })

    const hasActiveLink = computed(() => {
        return true
    })

    const hasChildren = computed(() => {
        return !!(item.value.items && item.value.items.length)
    })

    const route = useRoute();
    const isActiveLink = computed(() => {
        return item.value.link === route.path
    })

    function toggle() {
        collapsed.value = !collapsed.value
    }

    return {
        collapsed,
        isLink,
        isActiveLink,
        hasActiveLink,
        hasChildren,
        toggle
    }
}
