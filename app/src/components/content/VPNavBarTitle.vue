<script lang="ts" setup>

</script>

<template>
  <div class="VPNavBarTitle">
    <a class="title">
      <slot name="nav-bar-title-before"/>
      HDU-WIKI
      <slot name="nav-bar-title-after"/>
    </a>
  </div>
</template>

<style scoped>
.VPNavBarTitle {
  position: absolute;
  z-index: var(--vp-z-index-nav-title);
}

.title {
  display: flex;
  align-items: center;
  border-bottom: 1px solid transparent;
  width: 100%;
  height: var(--vp-nav-height);
  font-size: 16px;
  font-weight: 600;
  color: var(--vp-c-text-1);
  transition: opacity 0.25s;
}

@media (min-width: 960px) {
  .title {
    flex-shrink: 0;
  }

  .VPNavBarTitle.has-sidebar .title {
    border-bottom-color: var(--vp-c-divider);
  }
}

:deep(.logo) {
  margin-right: 8px;
  height: var(--vp-nav-logo-height);
}
</style>
