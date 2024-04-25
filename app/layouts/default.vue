<script setup>
const user = ref('blah');

const signOut = () => {
    user.value = undefined;
}
</script>

<template>
    <div class='layout'>
        <header>
            <NuxtLink class='link' to='/'>Home</NuxtLink>
            <div class='right-nav'>
                <NuxtLink class='link' to='/profile'>Profile</NuxtLink>
                <NuxtLink v-if="user" class='link' to='/' @click="signOut">Sign Out</NuxtLink>
                <NuxtLink v-else class='link' to='/login'>Sign In</NuxtLink>
            </div>
        </header>
        <slot />
    </div>
</template>

<style scoped>
header {
    position: sticky;
    display: flex;
    justify-content: space-between;
    margin: 0 16px;
    padding: 8px 0;
    font-size: 1.2em;
}

.layout {
    height: 100%;
}

.link {
    color: green;
    text-decoration: none;
    padding: 4px;
    border-radius: 4px;
}

.link:hover {
    background-color: lightcyan;
}

.right-nav {
    display: flex;
    gap: 16px;
}

p {
    color: white;
}
</style>
