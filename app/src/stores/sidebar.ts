import {computed, type ComputedRef, ref, type Ref, watchPostEffect} from 'vue'
import axios from "axios";
import {apiUrl} from "@/support/shared";

export interface Sidebar {
    text?: string
    link?: string
    items?: Sidebar[]
}

export interface SidebarControl {
    collapsed: Ref<boolean>
    collapsible: ComputedRef<boolean>
    isLink: ComputedRef<boolean>
    isActiveLink: Ref<boolean>
    hasActiveLink: ComputedRef<boolean>
    hasChildren: ComputedRef<boolean>

    toggle(): void
}

export function useSidebar(Loading: Ref<boolean>, sidebarGroups: Ref<Sidebar[]>) {
    axios.get(apiUrl + "/_meta").then((res) => {
        sidebarGroups.value = res.data.data;
        Loading.value = false;
    });
}

export function useSidebarControl(
    item: ComputedRef<Sidebar>
): SidebarControl {

    const collapsed = ref(false)

    const collapsible = computed(() => {
        return true
    })

    const isLink = computed(() => {
        return !!item.value.link
    })

    const isActiveLink = ref(false)


    const hasActiveLink = computed(() => {
        return !!isActiveLink.value;
    })

    const hasChildren = computed(() => {
        return !!(item.value.items && item.value.items.length)
    })

    watchPostEffect(() => {
        ;(isActiveLink.value || hasActiveLink.value) && (collapsed.value = false)
    })

    function toggle() {
        if (collapsible.value) {
            collapsed.value = !collapsed.value
        }
    }

    return {
        collapsed,
        collapsible,
        isLink,
        isActiveLink,
        hasActiveLink,
        hasChildren,
        toggle
    }
}