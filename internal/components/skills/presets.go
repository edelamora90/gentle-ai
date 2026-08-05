package skills

import "github.com/gentleman-programming/gentle-ai/v3/internal/model"

// sddSkills are the SDD orchestrator skills — always included.
var sddSkills = []model.SkillID{
	model.SkillSDDInit,
	model.SkillSDDExplore,
	model.SkillSDDResearch,
	model.SkillSDDPropose,
	model.SkillSDDSpec,
	model.SkillSDDDesign,
	model.SkillSDDTasks,
	model.SkillSDDApply,
	model.SkillSDDVerify,
	model.SkillSDDArchive,
	model.SkillSDDOnboard,
	model.SkillJudgmentDay,
}

// contributorSkills are this repository's own workflow skills. They stay
// selectable through the TUI skill picker and explicit `--skills` resolution,
// but no default preset installs them.
var contributorSkills = []model.SkillID{
	model.SkillGentleAIBench,
	model.SkillBranchPR,
	model.SkillIssueCreation,
	model.SkillCommentWriter,
	model.SkillRDDDefectWorkflow,
	model.SkillSystemicIssueTriage,
}

// designSkills back the sdd-visual phase. They ship with the recommended and
// full tiers, not with "minimal" — that preset is SDD-only by contract. In a
// minimal install sdd-visual still runs, but it works from its own phase
// guidance instead of delegating to these sub-skills.
var designSkills = []model.SkillID{
	model.SkillDesignResearcher,
	model.SkillFrontendDesign,
	model.SkillCopywritingForUI,
	model.SkillImageSourcingPolicy,
	model.SkillAccessibilityBaseline,
	model.SkillVisualCritic,
}

// selectableFoundationSkills is the canonical display order of every non-SDD
// skill. The TUI skill picker renders this order, so contributor skills keep
// their historical positions between the product skills; only preset
// membership changes.
var selectableFoundationSkills = []model.SkillID{
	model.SkillGoTesting,
	model.SkillGentleAIBench,
	model.SkillCreator,
	model.SkillImprover,
	model.SkillBranchPR,
	model.SkillIssueCreation,
	model.SkillSkillRegistry,
	model.SkillChainedPR,
	model.SkillCognitiveDoc,
	model.SkillCommentWriter,
	model.SkillWorkUnitCommits,
	model.SkillRDDDefectWorkflow,
	model.SkillSystemicIssueTriage,
}

// foundationSkills are the non-SDD product skills installed by the
// non-minimal presets: the selectable inventory minus the contributor skills.
var foundationSkills = excludeSkills(selectableFoundationSkills, contributorSkills)

func excludeSkills(src, exclude []model.SkillID) []model.SkillID {
	excluded := make(map[model.SkillID]struct{}, len(exclude))
	for _, id := range exclude {
		excluded[id] = struct{}{}
	}
	out := make([]model.SkillID, 0, len(src))
	for _, id := range src {
		if _, ok := excluded[id]; !ok {
			out = append(out, id)
		}
	}
	return out
}

// SkillsForPreset returns which skills should be installed for a given preset.
//
//   - "minimal" / PresetMinimal:       SDD skills only
//   - "ecosystem-only" / PresetEcosystemOnly: SDD + common framework skills
//   - "full-gentleman" / PresetFullGentleman: all available skills
//   - "custom" / PresetCustom:         empty (caller should provide explicit list)
func SkillsForPreset(preset model.PresetID) []model.SkillID {
	switch preset {
	case model.PresetMinimal:
		return concatSkills(sddSkills)
	case model.PresetEcosystemOnly:
		return concatSkills(sddSkills, designSkills, foundationSkills)
	case model.PresetFullGentleman:
		return concatSkills(sddSkills, designSkills, foundationSkills)
	case model.PresetCustom:
		return nil
	default:
		// Unknown preset — default to full.
		return concatSkills(sddSkills, designSkills, foundationSkills)
	}
}

// AllSkillIDs returns every selectable skill ID in canonical display order:
// the SDD suite first, then the non-SDD skills. It is the TUI picker's
// inventory and includes the contributor skills that no preset installs.
func AllSkillIDs() []model.SkillID {
	return concatSkills(sddSkills, designSkills, selectableFoundationSkills)
}

// concatSkills returns a fresh slice holding every group in order. It never
// appends into a package-level slice, so callers cannot alias shared backing
// arrays.
func concatSkills(groups ...[]model.SkillID) []model.SkillID {
	total := 0
	for _, group := range groups {
		total += len(group)
	}
	all := make([]model.SkillID, 0, total)
	for _, group := range groups {
		all = append(all, group...)
	}
	return all
}
