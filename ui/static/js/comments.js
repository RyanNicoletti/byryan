document.addEventListener("DOMContentLoaded", function () {
  const addCommentBtn = document.getElementById("add-comment-btn");
  const cancelCommentBtn = document.getElementById("cancel-comment-btn");
  const form = document.getElementById("add-comment-form");

  function toggleCommentForm() {
    form.classList.toggle("hidden");
  }

  if (addCommentBtn) {
    addCommentBtn.addEventListener("click", toggleCommentForm);
  }

  if (cancelCommentBtn) {
    cancelCommentBtn.addEventListener("click", toggleCommentForm);
  }
});
