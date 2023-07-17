import React from 'react';

import {
  EventBusSrv,
  EventBus,
  DashboardCursorSync,
  AnnotationEventUIModel,
  ThresholdsConfig,
  SplitOpen,
  CoreApp,
  DataFrame,
  InternalDataLinkSupplier,
} from '@grafana/data';

import { AdHocFilterItem } from '../Table/types';

import { SeriesVisibilityChangeMode } from '.';

/** @alpha */
export interface PanelContext {
  eventBus: EventBus;

  /** Dashboard panels sync */
  sync?: () => DashboardCursorSync;

  /** Information on what the outer container is */
  app?: CoreApp | 'string';

  /**
   * Called when a component wants to change the color for a series
   *
   * @alpha -- experimental
   */
  onSeriesColorChange?: (label: string, color: string) => void;

  onToggleSeriesVisibility?: (label: string, mode: SeriesVisibilityChangeMode) => void;

  canAddAnnotations?: () => boolean;
  canEditAnnotations?: (dashboardUID?: string) => boolean;
  canDeleteAnnotations?: (dashboardUID?: string) => boolean;
  onAnnotationCreate?: (annotation: AnnotationEventUIModel) => void;
  onAnnotationUpdate?: (annotation: AnnotationEventUIModel) => void;
  onAnnotationDelete?: (id: string) => void;

  /**
   * Used from visualizations like Table to add ad-hoc filters from cell values
   */
  onAddAdHocFilter?: (item: AdHocFilterItem) => void;

  /**
   * Enables modifying thresholds directly from the panel
   *
   * @alpha -- experimental
   */
  canEditThresholds?: boolean;

  /**
   * Shows threshold indicators on the right-hand side of the panel
   *
   * @alpha -- experimental
   */
  showThresholds?: boolean;

  /**
   * Called when a panel wants to change default thresholds configuration
   *
   * @alpha -- experimental
   */
  onThresholdsChange?: (thresholds: ThresholdsConfig) => void;

  /**
   * onSplitOpen is used in Explore to open the split view. It can be used in panels which has intercations and used in Explore as well.
   * For example TimeSeries panel.
   */
  onSplitOpen?: SplitOpen;

  /** For instance state that can be shared between panel & options UI  */
  instanceState?: any;

  /** Update instance state, this is only supported in dashboard panel context currently */
  onInstanceStateChange?: (state: any) => void;

  /**
   * Called when a panel is changing the sort order of the legends.
   */
  onToggleLegendSort?: (sortBy: string) => void;

  /**
   * Optional, only some contexts support this. This action can be cancelled by user which will result
   * in a the Promise resolving to a false value.
   */
  onUpdateData?: (frames: DataFrame[]) => Promise<boolean>;

  /**
   * Optional supplier for internal data links. If not provided a link pointing to Explore will be generated.
   * @internal
   */
  internalDataLinkSupplier?: InternalDataLinkSupplier;
}

export const PanelContextRoot = React.createContext<PanelContext>({
  eventBus: new EventBusSrv(),
});

/**
 * @alpha
 */
export const PanelContextProvider = PanelContextRoot.Provider;

/**
 * @alpha
 */
export const usePanelContext = () => React.useContext(PanelContextRoot);
