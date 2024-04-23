<template>
    <div class='page'>
        <h1 class='title'>Welcome to DJmocracy</h1>

        <div class='links'>
            <NuxtLink class='link' to='/party-search'>Find Party</NuxtLink>
            <NuxtLink class='link' to='/party-create'>Create Party</NuxtLink>
        </div>
    </div>
</template>

<style scoped>
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
