<script lang="ts" setup>
import {useWindowScroll} from '@vueuse/core'
import {computed, onMounted, ref} from 'vue'
import {useSidebarControl} from "@/stores/sidebar";
import VPIconAlignLeft from "@/components/icon/VPIconAlignLeft.vue";

defineProps<{
  open: boolean
}>()

defineEmits<{
  (e: 'open-menu'): void
}>()

const {isSidebarEnabled} = useSidebarControl()
// @ts-ignore
const {y} = useWindowScroll()

const navHeight = ref(0)

onMounted(() => {
  navHeight.value = parseInt(
      getComputedStyle(document.documentElement).getPropertyValue(
          '--vp-nav-height'
      )
  )
})

const empty = computed(() => {
  return false
})

const classes = computed(() => {
  return {
    VPLocalNav: true,
    fixed: empty.value,
    'reached-top': y.value >= navHeight.value
  }
})
</script>

<template>
  <div
      v-if="(!empty || y >= navHeight)"
      :class="classes"
  >
    <button
        v-if="isSidebarEnabled"
        :aria-expanded="open"
        aria-controls="VPSidebarNav"
        class="menu"
        @click="$emit('open-menu')"
    >
      <VPIconAlignLeft class="menu-icon"/>
      <span class="menu-text">
        {{ 'Menu' }}
      </span>
    </button>
  </div>
</template>

<style scoped>
.VPLocalNav {
  position: sticky;
  top: 0;
  /*rtl:ignore*/
  left: 0;
  z-index: var(--vp-z-index-local-nav);
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid var(--vp-c-gutter);
  border-bottom: 1px solid var(--vp-c-gutter);
  padding-top: 0;
  width: 100%;
  background-color: var(--vp-local-nav-bg-color);
}

.VPLocalNav.fixed {
  position: fixed;
}

.VPLocalNav.reached-top {
  border-top-color: transparent;
}

@media (min-width: 960px) {
  .VPLocalNav {
    display: none;
  }
}

.menu {
  display: flex;
  align-items: center;
  padding: 12px 24px 11px;
  line-height: 24px;
  font-size: 12px;
  font-weight: 500;
  color: var(--vp-c-text-2);
  transition: color 0.5s;
}

.menu:hover {
  color: var(--vp-c-text-1);
  transition: color 0.25s;
}

@media (min-width: 768px) {
  .menu {
    padding: 0 32px;
  }
}

.menu-icon {
  margin-right: 8px;
  width: 16px;
  height: 16px;
  fill: currentColor;
}

.VPOutlineDropdown {
  padding: 12px 24px 11px;
}

@media (min-width: 768px) {
  .VPOutlineDropdown {
    padding: 12px 32px 11px;
  }
}
</style>
