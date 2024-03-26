export default function Task() {
  function handleTitleClick() {
    alert("clicked Title");
  }

  function handleDescriptionClick(e) {
    let container = e.currentTarget;
    let input = document.createElement("input");
    input.type = "text";
    input.value = container.firstChild.textContent;
    container.replaceChild(input, container.firstChild);
    input.focus();
    input.onblur = handleDescriptionBlur;
    input.addEventListener("keyup", function (event) {
      if (event.key === "Enter") {
        handleDescriptionBlur(event);
      }
    });
  }

  function handleDescriptionBlur(e) {
    let container = e.currentTarget.parentNode;
    let value = container.firstChild.value;
    let textNode = document.createTextNode(value);
    container.replaceChild(textNode, container.firstChild);
  }

  return (
    <div className="Task">
      <div className="TaskTitle" onClick={handleTitleClick}>
        Task Title
      </div>
      <hr></hr>
      <div className="TaskDescription" onClick={handleDescriptionClick}>
        TASK DESCRIPTION
      </div>
    </div>
  );
}
