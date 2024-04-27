<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useUsersStore } from '@/stores/users';
import { watch } from 'vue';

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
        <header v-if="isLoggedIn">
            <button class='btn' @click="store.logOut">Logout</button>
        </header>
        <Suspense>
            <RouterView />
        </Suspense>
    </div>
</template>

<style scoped>
h1 {
    color: white;
}

.btn {
    padding: 6px;
    font-size: 1.2em;
    color: green;
    background-color: transparent;
    border-color: green;
    border-style: double;
    cursor: pointer;
}

.btn:hover {
    border-style: solid;
}

header {
    position: absolute;
    top: 8px;
    right: 8px;
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
