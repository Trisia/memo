<template>
  <split-pane :min="600" style="height: 100%;">
    <template v-slot:left>
      <div class="infinite-list-wrapper" style="overflow: auto">
        <ul v-infinite-scroll="load" class="list" :infinite-scroll-disabled="disabled">
          <li v-for="i in docArr" :key="i" class="list-item">{{ i }}</li>
          <p v-if="loading">Loading...</p>
          <p v-if="noMore">No more</p>
        </ul>
      </div>
    </template>
    <template v-slot:right>
      Right
    </template>
  </split-pane>

</template>

<script setup>
import SplitPane from '../components/SplitPane.vue';
import axios from 'axios';
import { computed, ref } from 'vue'

const docArr = ref([]);
const loading = ref(false)
const noMore = computed(() => docArr.value.length >= 20)
const disabled = computed(() => loading.value || noMore.value)
const load = () => {
  loading.value = true;
  axios.get("./api/doc/search?offset=0&limit=10").then(({ data }) => {
    console.log(data);
    for (let i = 0; i < data.length; i++) {
      docArr.value.push(data[i]);
    }
  }).catch((err) => {
    console.log(err);
  }).finally(() => {
    loading.value = false
  });
}
</script>

<style scoped>
.infinite-list-wrapper {
  height: 700px;
  text-align: center;
}

.infinite-list-wrapper .list {
  padding: 0;
  margin: 0;
  list-style: none;
}

.infinite-list-wrapper .list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  background: var(--el-color-danger-light-9);
  color: var(--el-color-danger);
}

.infinite-list-wrapper .list-item+.list-item {
  margin-top: 10px;
}
</style>