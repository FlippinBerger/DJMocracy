<script setup lang="ts">
import { RouterLink, RouterView, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useUsersStore } from '@/stores/users';
import Auth from './views/Auth.vue';
import { watch } from 'vue';

console.log('accessing the store in App view');

const store = useUsersStore();
const { isLoggedIn } = storeToRefs(store);

const router = useRouter();

watch(
    isLoggedIn,
    (logged) => {
        if (!logged) {
            router.replace('/auth');
        }
    },
    { immediate: true }
)
</script>

<template>
    <div class='page'>
        <Suspense>
            <RouterView />
        </Suspense>
    </div>
</template>

<style scoped>
h1 {
    color: white;
}

.page {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.title {
    margin: 16px;
}

.link {
    color: green;
    text-decoration: none;
    font-size: 1.5em;
    border: double;
    padding: 8px;
}

.link:hover {
    border: solid;
}

.links {
    display: flex;
    gap: 16px;
}
</style>
