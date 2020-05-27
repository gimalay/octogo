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
import com.example.todolist.model.HomeFilter
import com.example.todolist.model.HomeSorter
import com.example.todolist.ui.common.ActionsRow
import com.example.todolist.ui.common.DialogButton
import com.example.todolist.ui.common.TextFieldWithHint

@Composable
fun HomeHeader(
    onAddProject: (projectName: String) -> Unit,
    onApplyFilter: (HomeFilter) -> Unit,
    onApplySorter: (HomeSorter) -> Unit
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
    onApplySorter: (HomeSorter) -> Unit
) {
    val radioOptions = HomeSorter.projectFieldTitles
    val (selectedSorting, onSortingSelected) = +state { radioOptions[0] }
    val (isDescDirection, onDirectionChanged) = +state { HomeSorter.default.isDesc }

    DialogButton(
        text = "Sort",
        dialogTitle = "Sort by Project fields",
        onApply = {
            onApplySorter(
                HomeSorter(
                    field = HomeSorter.fieldByTitle[selectedSorting] ?: HomeSorter.default.field,
                    isDesc = isDescDirection
                )
            )
        }
    ) {
        Padding(padding = 24.dp) {
            Column {
                RadioGroup(
                    options = radioOptions,
                    selectedOption = selectedSorting,
                    onSelectedChange = onSortingSelected,
                    radioColor = (+MaterialTheme.colors()).primary
                )
                HeightSpacer(height = 8.dp)
                Row(
                    modifier = Spacing(left = 11.dp, top = 8.dp)
                ) {
                    Checkbox(
                        checked = isDescDirection,
                        onCheckedChange = onDirectionChanged,
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
    onApplyFilter: (HomeFilter) -> Unit
) {
    val (projectNameEditor, onProjectNameEditorChanged) = +state { EditorModel() }
    val (projectIdEditor, onProjectIdEditorChanged) = +state { EditorModel() }

    DialogButton(
        text = "Filter",
        dialogTitle = "Filter by Project fields",
        onApply = {
            onApplyFilter(
                HomeFilter(
                    projectId = projectIdEditor.text,
                    projectName = projectNameEditor.text
                )
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
        text = "Create",
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
