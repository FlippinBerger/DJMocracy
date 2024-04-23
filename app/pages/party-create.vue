<script setup>
const partyName = ref("");
const partyTime = ref(new Date());

const createParty = () => {
    // type of partyTime.value is a string
    console.log(`partyName: ${partyName.value} partyTime: ${partyTime.value} partyTime type: ${typeof (partyTime.value)}`);
}
</script>

<template>
    <div class='page'>
        <h1>Create Party</h1>
        <form class='form'>
            <input v-model="partyName" />
            <input type="datetime-local" v-model="partyTime" />
            <!-- TODO having it in here actually submits the form -->
            <!-- <button @click="createParty">Create</button> -->
        </form>
        <button class='btn' @click="createParty">Create</button>
    </div>
</template>

<style scoped>
.form {
    display: flex;
    gap: 8px;
}

.btn {
    margin-top: 16px;
    padding: 4px;
}
</style>
