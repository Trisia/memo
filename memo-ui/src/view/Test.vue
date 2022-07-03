<template>
  <div v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
    <div v-for="i in count" :key="i" >{{ i }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const count = ref(0)
const load = () => {
  count.value += 2
}
</script>

<style>
.infinite-list {
  height: 100%;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  height: 50px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}

</style>
