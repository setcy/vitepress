<script setup lang="ts">

import {RouterLink} from "vue-router";
</script>

<template>
  <nav>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/22">Content</RouterLink>
    <RouterLink to="/a/22">Admin</RouterLink>
  </nav>
</template>

<style scoped>

nav {
  font-size: 20px;
  text-align: center;
  border-bottom: 1px solid var(--color-border);
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

</style>
