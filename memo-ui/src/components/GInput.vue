<template>
    <div class="imformation">
        <input v-model="modelValue" :type="type" @focus="focus = true" @blur="focus = false"
            @input="emit('update:modelValue', modelValue)" :class="inputCss" autocomplete="off">
        <div :class="hitCss">{{ placeholder }}</div>
    </div>
</template>
<script setup>
import { ref, computed } from 'vue';

const props = defineProps({
    modelValue: String,
    placeholder: {
        type: String,
        default: '请输入',
    },
    type: {
        type: String,
        default: 'text',
    },
    size: {
        type: String,
        default: 'default',
    }
})

const focus = ref(false);
const emit = defineEmits(['update:modelValue'])

const inputCss = computed(() => {
    if (props.size == 'small') {
        return 'gg-input';
    } else {
        return 'gg-input-small';
    }
})

// 根据状态改变口令的提示消息位置和颜色
const hitCss = computed(() => {
    let res = {};
    if (props.size == 'small') {
        res['gg-div-small'] = true;
        if (focus.value) {
            res['gg-hit-focus-small'] = true
        }
        if (props.modelValue != '' && !focus.value) {
            res['gg-hit-focus-content-small'] = true
        }
    } else {
        res['gg-div'] = true;
        if (focus.value) {
            res['gg-hit-focus'] = true
        }
        if (props.modelValue != '' && !focus.value) {
            res['gg-hit-focus-content'] = true
        }
    }

    return res;
});
</script>
<style scoped>
.imformation {
    width: 100%;
    position: relative;
}

.gg-input {
    width: 100%;
    height: 56px;
    font-size: 16px;
    padding: 13px 15px 13px 15px;
    box-sizing: border-box;
    outline: none;
    border-radius: 8px;
    border: 1px solid #dadce0;
}

.gg-input:focus {
    border: 2px solid #1a73e8;
}

.gg-div {
    position: absolute;
    /*设置绝对定位， 那么给父级元素添加相对定位*/
    top: 17px;
    left: 12px;
    padding: 0 5px;
    font-size: 16px;
    font-weight: 400;
    font-family: roboto, 'Noto Sans Myanmar UI', arial, sans-serif;
    letter-spacing: 0.1px;
    /*字体间的空隙*/
    color: #5f6368;
    pointer-events: none;
    transition: all 0.15s ease;
}

.gg-hit-focus {
    top: -10px;
    background-color: #fff;
    font-size: 14px;
    color: #1a73e8;
}

.gg-hit-focus-content {
    top: -10px;
    background-color: #fff;
    font-size: 14px;
    color: #5f6368;
}


.gg-input-small {
    width: 100%;
    height: 56px;
    font-size: 16px;
    padding: 13px 15px 13px 15px;
    box-sizing: border-box;
    outline: none;
    border-radius: 8px;
    border: 1px solid #dadce0;
}

.gg-input-small:focus {
    border: 2px solid #1a73e8;
}

.gg-div-small {
    position: absolute;
    /*设置绝对定位， 那么给父级元素添加相对定位*/
    top: 17px;
    left: 12px;
    padding: 0 5px;
    font-size: 16px;
    font-weight: 400;
    font-family: roboto, 'Noto Sans Myanmar UI', arial, sans-serif;
    letter-spacing: 0.1px;
    /*字体间的空隙*/
    color: #5f6368;
    pointer-events: none;
    transition: all 0.15s ease;
}

.gg-hit-focus-small {
    top: -10px;
    background-color: #fff;
    font-size: 14px;
    color: #1a73e8;
}

.gg-hit-focus-content-small {
    top: -10px;
    background-color: #fff;
    font-size: 14px;
    color: #5f6368;
}
</style>