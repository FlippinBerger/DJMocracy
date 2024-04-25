<script setup>
    useHead({
        title: 'DJmocracy',
        meta: [
            {name: 'description', content: 'The amateur DJ and house party enthusiast\'s new best friend.'}
        ],
    })
</script>

<template>
    <div class='wrapper'>
        <NuxtLayout>
            <NuxtPage />
        </NuxtLayout>
    </div>
</template>

<style>
.wrapper {
    padding: 24px 24px;
    height: 100%;
}
</style>
