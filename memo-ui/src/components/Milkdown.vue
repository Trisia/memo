<template>
  <VueEditor :editor="editor" />
</template>

<script setup>
import 'katex/dist/katex.min.css';
import { Editor, rootCtx, defaultValueCtx } from "@milkdown/core";
import { nord } from "@milkdown/theme-nord";
import { VueEditor, useEditor } from "@milkdown/vue";
import { emoji } from "@milkdown/plugin-emoji";
import { history } from '@milkdown/plugin-history';
import { tooltip } from '@milkdown/plugin-tooltip';
import { menu, menuPlugin, defaultConfig } from "@milkdown/plugin-menu";
import { listener, listenerCtx } from '@milkdown/plugin-listener';
import { gfm } from '@milkdown/preset-gfm';

// const myMenu = menu.configure(menuPlugin, {
//   config: defaultConfig.map((section) => {
//     return section.map((item) => {
//       console.log(item);
//       if (item.type == 'select') {
//         item.text = '标题';
//         item.options = [
//           { id: '0', text: '正文' },
//           { id: '1', text: '标题 1' },
//           { id: '2', text: '标题 2' },
//           { id: '3', text: '标题 3' },
//           { id: '4', text: '标题 4' },
//         ];
//       }
//       switch (item.type) {
//         case 'select':

//           break;
//         case 'undo':
//           item.key = "撤销";
//           break;
//       }
//       return item;
//     });
//   }),
// });

const { editor } = useEditor((root) =>
  Editor.make()
    .config((ctx) => {
      ctx.set(rootCtx, root);
      ctx.set(defaultValueCtx, "# Milkdown 💖 Vue");
      ctx.get(listenerCtx).markdownUpdated((_, markdown) => {
        console.log(">> markdownUpdated", markdown);
      });
    })
    .use(gfm)
    .use(nord)
    .use(emoji)
    .use(history)
    .use(menu)
    .use(tooltip)
    .use(listener)
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