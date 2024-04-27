<script setup>
function getSongs() {
    return [
        {
            artist: 'Post Malone',
            title: 'Rockstar',
            votes: '3'
        },
        {
            artist: 'Taylor Swift',
            title: 'Fortnight',
            votes: '0'
        },
        {
            artist: 'Post Malone',
            title: 'Sunflower',
            votes: '-1'
        },
    ]
}
</script>

<template>
    <div class='party'>
        <h1>Welcome to party {{ $route.params.partyId }}</h1>
        <SongList :songs="getSongs()" />
    </div>
</template>

<style scoped>
.party {
    margin: 24px 16px;
}

li {
    color: white;
}
</style>
<style scoped>
li {
    color: white;
}
</style>
