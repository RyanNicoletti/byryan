document.addEventListener("DOMContentLoaded", function () {
  const addCommentBtn = document.getElementById("add-comment-btn");
  const form = document.getElementById("add-comment-form");

  if (form && form.dataset.hasErrors === "true") {
    form.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });
  }

  function toggleCommentForm() {
    form.classList.toggle("hidden");
  }

  if (addCommentBtn) {
    addCommentBtn.addEventListener("click", toggleCommentForm);
  }

  
});
