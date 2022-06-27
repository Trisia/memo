<template>
    <split-pane style="height:100%">
        <template v-slot:left>
            <split-pane>
                <template v-slot:left>
                    <div>AAA</div>
                </template>
                <template v-slot:right>
                    <div>BBB</div>
                </template>
            </split-pane>
        </template>
        <template v-slot:right>
            <div>CC</div>
        </template>
    </split-pane>
</template>
<script setup>
import SplitPane from '../components/SplitPane.vue'

</script>

<style>
html,
body {
    height: 100%;
}

#app {
    height: 100%;
}
</style>
