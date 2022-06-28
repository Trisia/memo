<script setup>
import { ref, reactive, computed } from 'vue'

const props = defineProps({
    layout: {
        type: String,
        default: '',
    },
    min: {
        type: Number,
        default: 150,
    },
    max: {
        type: Number,
        default: 0,
    }
})
const isVertical = computed(() => props.layout === 'vertical')

const container = ref()


const state = reactive({
    dragging: false,
    split: props.min
})

const boundSplit = computed(() => {
    const { split } = state;
    if (split < props.min) {
        return props.min;
    }

    if (props.max > 0 && split > props.max) {
        return props.max;
    }
    return split;
})

let startPosition = 0
let startSplit = 0

function dragStart(e) {
    state.dragging = true
    startPosition = isVertical.value ? e.pageY : e.pageX
    startSplit = boundSplit.value
}

function dragMove(e) {
    if (state.dragging) {
        const position = isVertical.value ? e.pageY : e.pageX
        const dp = position - startPosition
        state.split = startSplit + dp;
    }
}

function dragEnd() {
    state.dragging = false
}
</script>

<template>
    <div ref="container" class="split-pane" :class="{
        dragging: state.dragging,
        vertical: isVertical
    }" @mousemove="dragMove" @mouseup="dragEnd" @mouseleave="dragEnd">
        <div class="left" :style="{ [isVertical ? 'height' : 'width']: boundSplit + 'px' }">
            <slot name="left" />
            <div class="dragger" @mousedown.prevent="dragStart" />
        </div>
        <div class="right">
            <slot name="right" />
        </div>
    </div>
</template>

<style scoped>
.split-pane {
    display: flex;
    height: 100%;
    position: relative;
}

.split-pane.dragging {
    cursor: ew-resize;
}

.dragging .left,
.dragging .right {
    pointer-events: none;
}

.left {
    position: relative;
    height: 100%;
    border-right: 1px solid #D4D7DE;
}

.right {
    position: relative;
    height: 100%;
    width: 100%;
    flex: 1;
}


.left {
    border-right: 1px solid #D4D7DE;
}

.dragger {
    position: absolute;
    z-index: 3;
    top: 0;
    bottom: 0;
    right: -5px;
    width: 10px;
    cursor: ew-resize;
}

/* vertical */
@media (min-width: 721px) {
    .split-pane.vertical {
        display: block;
    }

    .split-pane.vertical.dragging {
        cursor: ns-resize;
    }

    .vertical .dragger {
        top: auto;
        height: 10px;
        width: 100%;
        left: 0;
        right: 0;
        bottom: -5px;
        cursor: ns-resize;
    }

    .vertical .left,
    .vertical .right {
        width: 100%;
    }

    .vertical .left {
        border-right: none;
        border-bottom: 1px solid var(--border);
    }
}

/* mobile */
@media (max-width: 720px) {

    .left,
    .right {
        width: 100% !important;
        height: 100% !important;
    }

    .dragger {
        display: none;
    }

    .split-pane .toggler {
        display: block;
    }

    .split-pane .right {
        display: none;
    }

    .split-pane.show-output .right {
        display: block;
    }

    .split-pane.show-output .left {
        display: none;
    }
}
</style>