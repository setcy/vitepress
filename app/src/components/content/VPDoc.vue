<script setup lang="ts">

import { Content } from '@/stores/content';

let rawHtml = `
<h2 class="toc-header" id="火车">
    火车
</h2>
<ul>
    <li>可在杭州火车东站（距学校较近）、杭州站（城站）乘车，地铁直达</li>
    <li>推荐网上购票，可以在师生办事大厅取票机取票或火车站取票机取票，带好身份证</li>
    <li>
        高铁动车大部分车次<strong>全价票</strong>
        支持刷身份证直接进出站，请注意购票时列车列表身份证标识
    </li>
</ul>
<h3 class="toc-header" id="火车优惠">
    火车优惠
</h3>
<ol>
    <li>在教务系统中填好火车优惠区间，优惠区间不可随意更改</li>
    <li>优惠区间只能在每年6/1-9/30以及当年12/1-次年3/31使用共四次（单程）</li>
    <li>购买学生票需在取票时带好带优惠磁条的学生证，否则无法取票</li>
    <li>学生票只享受半价硬座客票、加快票和空调票，动车二等座打75折，软座、软卧、动车一等票均无优惠</li>
</ol>`

</script>

<template>
  <div
    class="VPDoc"
    :class="{ 'has-sidebar': true, 'has-aside': true }"
  >
    <slot name="doc-top" />
    <div class="container">

      <div class="content">
        <div class="content-container">
          <slot name="doc-before" />
            <main class="main">
              <Content class="vp-doc" :htmlContent="$route.path + rawHtml" />
            </main>
          <slot name="doc-after" />
        </div>
      </div>
    </div>
    <slot name="doc-bottom" />
  </div>
</template>

<style scoped>
.VPDoc {
  padding: 32px 24px 96px;
  width: 100%;
}

.VPDoc .VPDocOutlineDropdown {
  display: none;
}

@media (min-width: 960px) and (max-width: 1279px) {
  .VPDoc .VPDocOutlineDropdown {
    display: block;
  }
}

@media (min-width: 768px) {
  .VPDoc {
    padding: 48px 32px 128px;
  }
}

@media (min-width: 960px) {
  .VPDoc {
    padding: 32px 32px 0;
  }

  .VPDoc:not(.has-sidebar) .container {
    display: flex;
    justify-content: center;
    max-width: 992px;
  }

  .VPDoc:not(.has-sidebar) .content {
    max-width: 752px;
  }
}

@media (min-width: 1280px) {
  .VPDoc .container {
    display: flex;
    justify-content: center;
  }

  .VPDoc .aside {
    display: block;
  }
}

@media (min-width: 1440px) {
  .VPDoc:not(.has-sidebar) .content {
    max-width: 784px;
  }

  .VPDoc:not(.has-sidebar) .container {
    max-width: 1104px;
  }
}

.container {
  margin: 0 auto;
  width: 100%;
}

.aside {
  position: relative;
  display: none;
  order: 2;
  flex-grow: 1;
  padding-left: 32px;
  width: 100%;
  max-width: 256px;
}

.left-aside {
  order: 1;
  padding-left: unset;
  padding-right: 32px;
}

.aside-container {
  position: fixed;
  top: 0;
  padding-top: calc(var(--vp-nav-height) + 32px);
  width: 224px;
  height: 100vh;
  overflow-x: hidden;
  overflow-y: auto;
  scrollbar-width: none;
}

.aside-container::-webkit-scrollbar {
  display: none;
}

.aside-curtain {
  position: fixed;
  bottom: 0;
  z-index: 10;
  width: 224px;
  height: 32px;
  background: linear-gradient(transparent, var(--vp-c-bg) 70%);
}

.aside-content {
  display: flex;
  flex-direction: column;
  min-height: calc(100vh - (var(--vp-nav-height) + 32px));
  padding-bottom: 32px;
}

.content {
  position: relative;
  margin: 0 auto;
  width: 100%;
}

@media (min-width: 960px) {
  .content {
    padding: 0 32px 128px;
  }
}

@media (min-width: 1280px) {
  .content {
    order: 1;
    margin: 0;
    min-width: 640px;
  }
}

.content-container {
  margin: 0 auto;
}

.VPDoc.has-aside .content-container {
  max-width: 688px;
}

.external-link-icon-enabled :is(.vp-doc a[href*='://'], .vp-doc a[target='_blank'])::after {
  content: '';
  color: currentColor;
}
</style>
