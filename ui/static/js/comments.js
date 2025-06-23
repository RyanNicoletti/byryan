addComment = document.getElementById("add-comment-btn");

function toggleCommentForm() {
  form = document.getElementById("add-comment-form");
  if (form.classList.contains("hidden")) {
    form.classList.remove("hidden");
  } else {
    form.classList.add("hidden");
  }
}

addComment.addEventListener("click", () => toggleCommentForm());
document.addEventListener("DOMContentLoaded", function () {
  const addCommentBtn = document.getElementById("add-comment-btn");
  const cancelCommentBtn = document.getElementById("cancel-comment-btn");
  const form = document.getElementById("add-comment-form");

  function toggleCommentForm() {
    if (form.classList.contains("hidden")) {
      form.classList.remove("hidden");
      addCommentBtn.style.display = "none";
    } else {
      form.classList.add("hidden");
      addCommentBtn.style.display = "block";
    }
  }

  if (addCommentBtn) {
    addCommentBtn.addEventListener("click", toggleCommentForm);
  }

  if (cancelCommentBtn) {
    cancelCommentBtn.addEventListener("click", toggleCommentForm);
  }
});
