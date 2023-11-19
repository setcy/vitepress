import type {Ref} from 'vue';
import {defineComponent, h} from 'vue';
import axios from "axios";
import {apiUrl} from "@/support/shared";

export interface Aside {
    title?: string
    anchor?: string
    children?: Aside[]
}

export function useContent(path: Ref<string>, Loading: Ref<boolean>, content: Ref<string>, aside: Ref<any>) {
    axios.get(apiUrl + "/_content" + path.value).then((res) => {
        content.value = res.data.data.content;
        aside.value = res.data.data.toc;
        Loading.value = false;
    });
}

export const Content = defineComponent({
    name: 'VitePressContent',
    props: {
        rawHtml: {
            type: String,
            required: true
        },
        loading: {
            type: Boolean,
            required: true
        }
    },
    setup(props) {
        return () => props.loading
            ? h('div', {class: 'spinner'})  // Loading
            : h('div', {
                innerHTML: props.rawHtml
            });
    }

});