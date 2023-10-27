<template>
  <v-app>
    <v-navigation-drawer permanent>
      <v-list>
        <v-list-item title="Jarvis" subtitle="Welcome to this page">
        </v-list-item>
      </v-list>

      <v-divider></v-divider>

      <v-list :lines="false" density="compact" nav>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :value="item"
          color="blue">
          <template v-slot:prepend>
            <v-icon :icon="item.icon"></v-icon>
          </template>

          <v-list-item-title>{{ item.text }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container fluid> text </v-container>
    </v-main>
    <v-footer app
      ><v-row class="d-flex align-center">
        <v-col cols="1" class="d-flex justify-end">
          <v-btn class="rounded-circle" icon="mdi-microphone"></v-btn>
        </v-col>
        <v-col cols="10">
          <v-textarea rows="1" hide-details="auto"></v-textarea>
        </v-col>
        <v-col cols="1">
          <v-btn class="rounded-circle" icon="mdi-send"></v-btn>
        </v-col>
      </v-row>
    </v-footer>
  </v-app>
</template>

<script setup>
import { ref } from "vue";
const items = ref([
  { text: "My Files", icon: "mdi-folder" },
  { text: "Shared with me", icon: "mdi-account-multiple" },
  { text: "Starred", icon: "mdi-star" },
  { text: "Recent", icon: "mdi-history" },
  { text: "Offline", icon: "mdi-check-circle" },
  { text: "Uploads", icon: "mdi-upload" },
  { text: "Backups", icon: "mdi-cloud-upload" },
]);
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
