import piniaPluginPersistedState from "pinia-plugin-persistedstate"

export default defineNuxtPlugin({
    name: 'pinia-persist',
    async setup (nuxtApp) {
        const pinia = usePinia();

        pinia.use(piniaPluginPersistedState);
    }
})

