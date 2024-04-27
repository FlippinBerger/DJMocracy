<script setup>
    import {ref} from 'vue';
    const {data: stuff} = await fetch('http://localhost:1323/');
    const s = ref(stuff);
</script>

<template>
    <div class='page'>
        <h1 class='title'>Welcome to DJmocracy</h1>

        <div class='links'>
            <RouterLink class='link' to='/party-search'>Find Party</RouterLink>
            <RouterLink class='link' to='/party-create'>Create Party</RouterLink>
        </div>
            <RouterLink class='link' to='/parties/5'>Party</RouterLink>

        <h1 v-if="s">{{s}}</h1>
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
