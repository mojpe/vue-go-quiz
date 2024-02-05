<!-- src/components/Quiz.vue -->
<template>
    <div>
      <Question
        v-if="questions.length > 0"
        :question="questions[currentIndex]"
        :currentIndex="currentIndex"
        :totalQuestions="totalQuestions"
        @changeQuestion="handleQuestionChange"
      />
      <div v-else>Loading questions...</div>
    </div>
  </template>
  
  <script setup lang="ts">
    import { ref, onMounted } from 'vue';
    import Question from '@/components/Question.vue'; // Adjust the path accordingly
    import axios from 'axios';
    
    const questions = ref([]);
    const currentIndex = ref(0);
    const totalQuestions = ref(questions.length);
    
    onMounted(fetchQuizQuestions);
    
    async function fetchQuizQuestions() {
        try {
            const response = await axios.get('http://localhost:8081/api/questions');
            questions.value = response.data;
            totalQuestions.value = questions.value.length;
        } catch (error) {
            console.error('Error fetching quiz questions:', error);
        }
    }
    
    function handleQuestionChange(newIndex) {
        currentIndex.value = newIndex;
    }
  </script>
  
  <style scoped>
  /* Add any additional custom styles here */
  </style>
  