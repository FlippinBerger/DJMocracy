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

p {
    color: white;
}
</style>
