package com.example.todolist.ui.home

import androidx.compose.Composable
import androidx.compose.state
import androidx.compose.unaryPlus
import androidx.ui.core.EditorModel
import androidx.ui.core.Text
import androidx.ui.core.dp
import androidx.ui.layout.*
import androidx.ui.material.Checkbox
import androidx.ui.material.MaterialTheme
import androidx.ui.material.RadioGroup
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.model.ViewModelOuterClass.SortDirection
import com.example.todolist.ui.common.ActionsRow
import com.example.todolist.ui.common.DialogButton
import com.example.todolist.ui.common.TextFieldWithHint

// todo: remove this helper when RadioBox will be able to keep objects instead of only strings
object ProjectSortHelper {
    private val projectFields = listOf(
        ViewModel.Home.Project.ID_FIELD_NUMBER,
        ViewModel.Home.Project.NAME_FIELD_NUMBER
    )

    val projectFieldTitles = listOf("Id", "Name")
    val fieldByTitle: Map<String, Int> = (projectFieldTitles zip projectFields).toMap()
    val titleByField: Map<Int, String> = (projectFields zip projectFieldTitles).toMap()
}

@Composable
fun HomeHeader(
    onAddProject: (projectName: String) -> Unit,
    onApplyFilter: (Location.Home.Filter) -> Unit,
    onApplySorter: (Location.Home.Sorter) -> Unit
) {
    Column {
        ActionsRow {
            SorterOfProjectsAction(onApplySorter = onApplySorter)
            WidthSpacer(width = 12.dp)
            FilterOfProjectsAction(onApplyFilter = onApplyFilter)
            WidthSpacer(width = 12.dp)
            NewProjectAction(onAddProject = onAddProject)
        }
    }
}


@Composable
fun SorterOfProjectsAction(
    onApplySorter: (Location.Home.Sorter) -> Unit
) {
    val sortOptions = ProjectSortHelper.projectFieldTitles
    val (sorting, setSorting) = +state {
        ProjectSortHelper.titleByField[ViewModel.Home.Project.NAME_FIELD_NUMBER]
    }
    val (direction, setDirection) = +state { SortDirection.Desc }

    DialogButton(
        buttonText = "Sort",
        dialogTitle = "Sort by Project fields",
        onApply = {
            onApplySorter(
                Location.Home.Sorter
                    .newBuilder()
                    .setProjectField(
                        ProjectSortHelper.fieldByTitle[sorting] ?: ViewModel.Home.Project.NAME_FIELD_NUMBER
                    )
                    .setDirection(direction)
                    .build()
            )
        }
    ) {
        Padding(padding = 24.dp) {
            Column {
                RadioGroup(
                    options = sortOptions,
                    selectedOption = sorting,
                    onSelectedChange = setSorting,
                    radioColor = (+MaterialTheme.colors()).primary
                )
                HeightSpacer(height = 8.dp)
                Row(
                    modifier = Spacing(left = 11.dp, top = 8.dp)
                ) {
                    Checkbox(
                        checked = direction == SortDirection.Desc,
                        onCheckedChange = {
                            setDirection(if (it) SortDirection.Desc else SortDirection.Asc)
                        },
                        color = (+MaterialTheme.colors()).primaryVariant
                    )
                    WidthSpacer(width = 14.dp)
                    Padding(padding = 4.dp) {
                        Text("Descending sorting")
                    }
                }

            }
        }
    }
}

@Composable
fun FilterOfProjectsAction(
    onApplyFilter: (Location.Home.Filter) -> Unit
) {
    val (projectNameEditor, onProjectNameEditorChanged) = +state { EditorModel() }
    val (projectIdEditor, onProjectIdEditorChanged) = +state { EditorModel() }

    DialogButton(
        buttonText = "Filter",
        dialogTitle = "Filter by Project fields",
        onApply = {
            onApplyFilter(
                Location.Home.Filter
                    .newBuilder()
                    .setProjectId(projectIdEditor.text)
                    .setProjectName(projectNameEditor.text)
                    .build()
            )
        }
    ) {
        Padding(padding = 24.dp) {
            TextFieldWithHint(
                label = "ID:",
                hint = "Please, enter search keyword",
                value = projectIdEditor,
                onChange = onProjectIdEditorChanged
            )
        }
        Padding(padding = 24.dp) {
            TextFieldWithHint(
                label = "Name:",
                hint = "Please, enter search keyword",
                value = projectNameEditor,
                onChange = onProjectNameEditorChanged
            )
        }
    }
}

@Composable
fun NewProjectAction(onAddProject: (String) -> Unit) {
    val (projectNameEditor, onProjectNameEditorChanged) = +state { EditorModel() }

    DialogButton(
        buttonText = "Create",
        disabled = projectNameEditor.text.isEmpty(),
        dialogTitle = "Create Project",
        onApply = {
            onAddProject(projectNameEditor.text)
            onProjectNameEditorChanged(EditorModel(""))
        }
    ) {
        Padding(padding = 24.dp) {
            TextFieldWithHint(
                label = "Name:",
                value = projectNameEditor,
                hint = "Please, enter name",
                onChange = onProjectNameEditorChanged
            )
        }
    }
}
