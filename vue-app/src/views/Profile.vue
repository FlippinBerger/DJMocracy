<script setup>
</script>

<template>
    <div class='page'>
        <h1 class='title'>Profile page</h1>

        <!-- TODO fetch all the users parties here -->
    </div>
</template>

<style scoped>
.title {
    margin: 16px;
}

.controls {
    display: flex;
    flex-direction: column;
    /* justify-content: center; */
    width: 75%;
    margin: 0 auto;
    gap: 8px
}

input {
    padding: 4px;
    min-width: 20%;
}

.btn {
    padding: 4px;
    /* width: 10%; */
    align-self: center;
    min-width: 25%;
}

@media (min-width: 1024px) {
    .controls {
        flex-direction: row;
        justify-content: center;
    }
    .btn {
        min-width: 5%;
        width: 10%;
    }
}

</style>
