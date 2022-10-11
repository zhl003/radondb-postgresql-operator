/*
Copyright 2021 RadonDB.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import "sigs.k8s.io/controller-runtime/pkg/reconcile"

// updateReconcileResult creates a new Result based on the new and existing results provided to it.
// This includes setting "Requeue" to true in the Result if set to true in the new Result but not
// in the existing Result, while also updating RequeueAfter if the RequeueAfter value for the new
// result is less the the RequeueAfter value for the existing Result.
func updateReconcileResult(currResult, newResult reconcile.Result) reconcile.Result {

	if newResult.Requeue {
		currResult.Requeue = true
	}

	if newResult.RequeueAfter != 0 {
		if currResult.RequeueAfter == 0 || newResult.RequeueAfter < currResult.RequeueAfter {
			currResult.RequeueAfter = newResult.RequeueAfter
		}
	}

	return currResult
}
