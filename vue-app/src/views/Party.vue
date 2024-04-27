<script setup>
import SongList from '@/components/SongList.vue'
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
    <!-- <div class='page'> -->
    <div class='party'>
        <h1 class='title'>Welcome to party {{ $route.params.partyId }}</h1>
        <div class='list-wrapper'>
            <SongList :songs="getSongs()" />
        </div>
    </div>
    <!-- </div> -->
</template>

<style scoped>
.party {
    margin: 24px 16px;
    align-self: flex-start;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
}

.list-wrapper {
    margin: 0 48px;
}

.title {
    align-self: center;
}

</style>

