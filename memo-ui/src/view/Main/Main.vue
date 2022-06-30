<template>
    <split-pane :min="600" style="height:100%">
        <template v-slot:left>
            <split-pane :min="250">
                <template v-slot:left>
                    <div style="height: 100%; background-color: #f3f4f6; padding: 15px;" v-loading="userArealoading">
                        <div style="padding: 30px 15px 30px 15px;text-align: center;">
                            <img style=" border-radius: 50%;" width="90" height="90" :src="userInfo.avatar" alt="">
                            <div style="margin: 5px;font-weight: 600; font-size: 18px">{{ userInfo.username }}</div>
                        </div>
                        <div>
                            <h3 style="margin-left: 5px;">专栏</h3>
                            <div style="background-color: #fff; border-radius: 3px;" class="doc-tag-item"
                                :class="{ 'doc-tag-active': selectedTag.id == 0 }" @click="handleSelectTag({ id: 0 })">
                                <div>全部</div>
                            </div>
                            <div
                                style="margin-top: 20px;background-color: #fff; border-radius: 3px; padding:10px 10px 24px 10px;">
                                <div class="doc-tag-item" v-for="item in docTagArr" @click="handleSelectTag(item)"
                                    :class="{ 'doc-tag-active': selectedTag.id == item.id }">
                                    <div>{{ item.name }}</div>
                                    <div>{{ item.count }}</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </template>
                <template v-slot:right>
                    <div style="padding: 15px">
                        <el-input v-model="docParam.keyword" @keydown.enter="handlerSearchDoc" :prefix-icon="Search"
                            placeholder="搜索您的内容"></el-input>
                    </div>
                    <div style="padding: 10px 10px 0 10px;height: 100%; overflow: auto;"
                        v-infinite-scroll="handlerDocLAppendLoad" v-loading="docListArealoading">
                        <div v-for="item in docArr" :key="item.id" class="doc-item"
                            :class="{ 'doc-item-active': item.id == selectedDocId }" @click="handleSelectDoc(item)">
                            doc
                        </div>
                        <div v-if="docPullNewLoading" style="margin-top:45px" v-loading="docPullNewLoading"> </div>
                        <div style="margin-top:15px; text-align: center;">已经到底了</div>
                    </div>
                </template>
            </split-pane>
        </template>
        <template v-slot:right>
            <div>CC</div>
        </template>
    </split-pane>
</template>
<script setup>
import axios from 'axios';
import { onMounted, reactive, ref } from 'vue';
import SplitPane from '../../components/SplitPane.vue'
import avatar from '../../assets/avatar.png'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'

const userArealoading = ref(false);
const docListArealoading = ref(false);
const docPullNewLoading = ref(false);
const noMoreDoc = ref(false);

const selectedDocId = ref(0);

// 已经被选中的文档标签
const selectedTag = reactive({ id: 0, name: "", count: 0 })
const docParam = reactive({
    keyword: '',
    limit: 10,
    offset: 0
});

var userInfo = reactive({
    "id": 0,
    "createdAt": "",
    "updatedAt": "",
    "username": "",
    "email": "",
    "typ": 0,
    "avatar": avatar
});

const docTagArr = ref([]);
const docArr = ref([]);

onMounted(() => {
    userArealoading.value = true;
    getUserInfo();
    getTags();
    docListArealoading.value = true;
    handlerSearchDoc();
});


// 查询用户信息
const getUserInfo = () => {
    axios.get("./api/user/info").then(({ data }) => {
        for (const key in data) {
            let value = data[key]
            if (key == "avatar" && value) {
                if (value[0] == '/') {
                    // '/' : jpg
                    value = "data:image/jpeg;base64," + value
                } else {
                    // 'i' : png
                    value = "data:image/png;base64," + value
                }
            }
            userInfo[key] = value;
        }
    }).catch((err) => {
        ElMessage.error(err.response.data);
    });
}

const getTags = () => {
    axios.get("./api/tag/list").then(({ data }) => {
        docTagArr.value = data;
    }).catch((err) => {
        ElMessage.error(err.response.data)
    }).finally(() => {
        userArealoading.value = false;
    })
}

// 选中Tag id为0 表示全部
const handleSelectTag = (item) => {
    if (item.id == 0) {
        selectedTag.id = 0;
        selectedTag.name = "";
        selectedTag.count = 0;
    } else {
        for (const key in item) {
            selectedTag[key] = item[key];
        }
    }
}


// 搜索文档
const handlerSearchDoc = () => {
    let param = `?offset=${docParam.offset}&limit=${docParam.limit}`;
    if (docParam.keyword) {
        param += "&keyword=" + docParam.keyword.trim();
    }
    if (selectedTag.id > 0) {
        param += "&tagId=" + selectedTag.id;

    }
    axios.get("./api/doc/search" + param).then(({ data }) => {
        docArr.value = data;
    }).catch((err) => {
        ElMessage.error(err.response.data)
    }).finally(() => {
        docListArealoading.value = false;
    })
}

const handlerDocLAppendLoad = () => {
    if (noMoreDoc.value) {
        // 已经到底了
        return;
    }

    docPullNewLoading.value = true;
    docParam.offset += docParam.limit;
    let param = `?offset=${docParam.offset}&limit=${docParam.limit}`;
    if (docParam.keyword) {
        param += "&keyword=" + docParam.keyword.trim();
    }
    if (selectedTag.id > 0) {
        param += "&tagId=" + selectedTag.id;

    }
    axios.get("./api/doc/search" + param).then(({ data }) => {
        for (let i = 0; i < data.length; i++) {
            docArr.push(data[i]);
        }
        if (data.length < docParam.limit) {
            noMoreDoc.value = true;
        }
    }).catch((err) => {
        ElMessage.error(err.response.data)
    }).finally(() => {
        docPullNewLoading.value = false;
    })
}


// 选中文档时
const handleSelectDoc = (item) => {
    //TODO: 文档变更
    selectedDocId.value = item.id;
}

</script>

<style>
html,
body {
    height: 100%;
}

#app {
    height: 100%;
}


.doc-tag-item {
    display: flex;
    padding: 10px 15px;
    justify-content: space-between;
    color: #606266;
    user-select: none;
}

.doc-tag-item:hover {
    color: #50a6ff;
    cursor: pointer;
    font-weight: 600;
}

.doc-tag-active {
    color: #50a6ff;
    font-weight: 600;
}

.doc-item {
    border-bottom: 1px solid #E4E7ED;
    padding: 10px
}

.doc-item:hover {
    background-color: #f3f4f6;
}

.doc-item-active {
    background-color: #f3f4f6;
    color: #50a6ff;
}
</style>
