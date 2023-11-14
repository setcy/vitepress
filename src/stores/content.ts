import { defineComponent, h } from 'vue';

export const Content = defineComponent({
    name: 'VitePressContent',
    props: {
        htmlContent: {
            type: String,
            required: true
        }
    },
    render() {
        return h('div', {
            innerHTML: this.htmlContent
        });
    }
});