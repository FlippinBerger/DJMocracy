<template>
    <div class='layout'>
        <header>
            <NuxtLink class='link' to='/'>Home</NuxtLink>
            <NuxtLink class='link' to='/profile'>Profile</NuxtLink> 
        </header>
        <slot />
    </div>
</template>

<style scoped>
header {
    display: flex;
    justify-content: space-between;
    margin: 0 16px;
    font-size: 1.2em;
}

.layout {
    height: 100%;
}

.link {
    color: green;
    text-decoration: none;
}

p {
    color: white;
}
</style>
