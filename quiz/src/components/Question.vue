<!-- src/components/Question.vue -->
<template>
    <div class="container mt-5" v-if="question && question.title">
      <div class="d-flex justify-content-center row">
        <div class="col-md-10 col-lg-10">
          <div class="border">
            <div class="question bg-white p-3 border-bottom">
              <div class="d-flex flex-row justify-content-between align-items-center mcq">
                <h4>MCQ Quiz</h4><span>({{ currentIndex + 1 }} of {{ totalQuestions }})</span>
              </div>
            </div>
            <div class="question bg-white p-3 border-bottom">
              <div class="d-flex flex-row align-items-center question-title">
                <h3 class="text-danger">Q.</h3>
                <h5 class="mt-1 ml-2">{{ question.title }}</h5>
              </div>
              <div class="ans ml-2">
                <template v-for="(answer, answerIndex) in question.answers" :key="answerIndex">
                  <div class="form-check">
                    <input class="form-check-input" type="radio" :name="`question_${currentIndex}`" :value="answer" v-model="selectedAnswer">
                    <label class="form-check-label">{{ answer }}</label>
                  </div>
                </template>
              </div>
            </div>
            <div class="d-flex flex-row justify-content-between align-items-center p-3 bg-white">
              <button class="btn btn-primary" type="button" @click="previousQuestion">
                <i class="fa fa-angle-left mr-1"></i>Previous
              </button>
              <button class="btn btn-primary" type="button" @click="nextQuestion">
                Next<i class="fa fa-angle-right ml-2"></i>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { defineProps, ref, watchEffect, getCurrentInstance } from 'vue';
  
  const props = defineProps(['question', 'currentIndex', 'totalQuestions']);
  const selectedAnswer = ref('');
  const instance = getCurrentInstance();
  
  // Watch for changes in the 'question' prop
  watchEffect(() => {
    if (props.question) {
      selectedAnswer.value = ''; // Reset selected answer when the question changes
      // You can perform additional actions when the 'question' prop changes
    }
  });
  
  function previousQuestion() {
    // Handle logic to move to the previous question
    console.log('Previous question');
    if (props.currentIndex > 0) {
        instance.emit('changeQuestion', props.currentIndex - 1);
    }
  }
  
  function nextQuestion() {
    // Handle logic to move to the next question
    console.log('Next question');
    if (props.currentIndex < props.totalQuestions - 1) {
        instance.emit('changeQuestion', props.currentIndex + 1);
    }
  }
  </script>
  
  <style scoped>
  /* Add any additional custom styles here */
  </style>
  