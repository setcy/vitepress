import {defineComponent, h, ref, watch} from 'vue';
import axios from "axios";

export const Content = defineComponent({
    name: 'VitePressContent',
    props: {
        path: {
            type: String,
            required: true
        }
    },
    setup(props) {
        const isLoading = ref(true);
        const content = ref('');

        const fetchData = (path : string) => {
            axios.get("http://localhost:8080/content" + path).then((res) => {
                content.value = res.data.data.content;
                isLoading.value = false;
            });
        };

        fetchData(props.path);

        watch(() => props.path, (newPath) => {
            fetchData(newPath);
        });

        return () => isLoading.value
            ? h('div', { class: 'spinner' })  // Loading
            : h('div', {
                innerHTML: content.value
            });
    }

});