<script lang="ts" setup>
import { useScrollLock } from '@vueuse/core'
import { ref, watch } from 'vue'
import VPSidebarItem from './VPSidebarItem.vue'
import {inBrowser} from "@/support/shared";

let sidebarGroups = [
  {
    text: 'Home',
    children: [
      {
        text: 'Child',
        link: '/child',
      }
    ],
  },
  {
    text: 'About',
    link: '/about',
    children: [],
  },
  {
    text: 'Contact',
    link: '/contact',
    children: [],
  },
]

const props = defineProps<{
  open: boolean
  hasSidebar: boolean
}>()

// a11y: focus Nav element when menu has opened
const navEl = ref<HTMLElement | null>(null)
const isLocked = useScrollLock(inBrowser ? document.body : null)

watch(
    [props, navEl],
    () => {
      if (props.open) {
        isLocked.value = true
        navEl.value?.focus()
      } else isLocked.value = false
    },
    { immediate: true, flush: 'post' }
)
</script>

<template>
  <aside
      v-if="hasSidebar"
      class="VPSidebar"
      :class="{ open }"
      ref="navEl"
      @click.stop
  >
    <div class="curtain" />

    <nav class="nav" id="VPSidebarNav" aria-labelledby="sidebar-aria-label" tabindex="-1">
      <span class="visually-hidden" id="sidebar-aria-label">
        Sidebar Navigation
      </span>

      <slot name="sidebar-nav-before" />

      <div v-for="item in sidebarGroups" :key="item.text" class="group">
        <VPSidebarItem :item="item" :depth="0" />
      </div>

      <slot name="sidebar-nav-after" />
    </nav>
  </aside>
</template>

<style scoped>
.VPSidebar {
  position: fixed;
  top: var(--vp-layout-top-height, 0px);
  bottom: 0;
  left: 0;
  z-index: var(--vp-z-index-sidebar);
  padding: 32px 32px 96px;
  width: calc(100vw - 64px);
  max-width: 320px;
  background-color: var(--vp-sidebar-bg-color);
  opacity: 0;
  box-shadow: var(--vp-c-shadow-3);
  overflow-x: hidden;
  overflow-y: auto;
  transform: translateX(-100%);
  transition: opacity 0.5s, transform 0.25s ease;
  overscroll-behavior: contain;
}

.VPSidebar.open {
  opacity: 1;
  visibility: visible;
  transform: translateX(0);
  transition: opacity 0.25s,
  transform 0.5s cubic-bezier(0.19, 1, 0.22, 1);
}

.dark .VPSidebar {
  box-shadow: var(--vp-shadow-1);
}

.nav {
  outline: 0;
}

.group + .group {
  border-top: 1px solid var(--vp-c-divider);
  padding-top: 10px;
}
</style>
