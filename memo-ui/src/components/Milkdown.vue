<template>
    <VueEditor :editor="editor" />
</template>

<script setup>
import { Editor, rootCtx, defaultValueCtx } from "@milkdown/core";
import { nord } from "@milkdown/theme-nord";
import { VueEditor, useEditor } from "@milkdown/vue";
import { commonmark } from "@milkdown/preset-commonmark";
import { emoji } from "@milkdown/plugin-emoji";
import { history } from '@milkdown/plugin-history';
import { tooltip } from '@milkdown/plugin-tooltip';
import { menu } from "@milkdown/plugin-menu";


const { editor } = useEditor((root) =>
    Editor.make()
        .config((ctx) => {
            ctx.set(rootCtx, root);
            ctx.set(defaultValueCtx, "# Milkdown 💖 Vue");
        })
        .use(nord)
        .use(emoji)
        .use(commonmark)
        .use(history)
        .use(menu)
        .use(tooltip)
);

</script>
<style>
@font-face {
  font-family: 'Material Icons Outlined';
  font-style: normal;
  font-weight: 400;
  src: url(./gok-H7zzDkdnRel8-DQ6KAXJ69wP1tGnf4ZGhUce.woff2) format('woff2');
}

.material-icons-outlined {
  font-family: 'Material Icons Outlined';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  white-space: nowrap;
  word-wrap: normal;
  direction: ltr;
  -webkit-font-feature-settings: 'liga';
  -webkit-font-smoothing: antialiased;
}
</style>