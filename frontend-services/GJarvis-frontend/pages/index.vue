<template>
    <div>
        <v-toolbar dark prominent image="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg">
            <v-app-bar-nav-icon></v-app-bar-nav-icon>

            <v-toolbar-title>Vuetify</v-toolbar-title>

            <v-spacer></v-spacer>

            <v-btn icon>
                <v-icon>mdi-export</v-icon>
            </v-btn>
        </v-toolbar>
    </div>
    <div>
        <h1>aaa{{ count }}</h1>
        <h1>12312321312</h1>
    </div>
    <v-sheet :height="200" :width="200"></v-sheet>
</template>

<script setup>
const { data: count } = await useFetch('/api/v1/voice-assistance-service/hello')
</script>



<!-- <script lang="ts" setup>
import { ref, onMounted } from "vue";
import axios from "axios";

const message = ref('');

onMounted(async () => {
    // try {
    //     const response = 
    await axios.get("http://localhost:8080/api/v1/voice-assistance-service/hello")
        .then((response) => message.value = response.data)
        .catch((error) => console.log(error))
    //     message.value = response.data;
    // } catch (error) {
    //     console.log(error);
    // }
}); -->
<!-- </script> -->
